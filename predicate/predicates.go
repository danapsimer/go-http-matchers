// Package predicate defines an interface named Predicate that has a function named Accept that returns true or false
// based on the value passed to it.  The package also defines a set of predicates useful for making decisions about
// http.Request objects.  For instance, the QueryParameterEquals(name, value string) will return a Predicate that
// expects a *http.Request and will return true if the value of query parameter named 'name' equals 'value'.
package predicate

// Licensed to BlueSoft Development, LLC under one or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information regarding copyright ownership.  BlueSoft Development, LLC
// licenses this file to you under the Apache License, Version 2.0 (the "License"); you may not use this file except in
// compliance with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations under the License.

import (
	. "github.com/bluesoftdev/go-http-matchers/extractor"
	"strings"
)

// Predicate is a class that can accept or reject a value based on some condition.
type Predicate interface {
	Accept(interface{}) bool
}

// PredicateFunc is an implementation of Predicate that is a function and calls itself on a call to Accept
type PredicateFunc func(interface{}) bool

// PredicateFunc.Accept calls the predicate func with the passed value.
func (pf PredicateFunc) Accept(v interface{}) bool {
	return pf(v)
}

// And returns a predicate that is true if all of the passed predicate are true for the input.  Furthermore, it stops
// executing predicates after the first false one.
func And(predicates ...Predicate) Predicate {
	return PredicateFunc(func(v interface{}) bool {
		for _, p := range predicates {
			if !p.Accept(v) {
				return false
			}
		}
		return true
	})
}

// Or returns a predicate that is true if any of the passed predicate are true.  Furthermore, it stops executing
// predicates after the first true one.
func Or(predicates ...Predicate) Predicate {
	return PredicateFunc(func(v interface{}) bool {
		for _, p := range predicates {
			if p.Accept(v) {
				return true
			}
		}
		return false
	})
}

// Not returns a predicate that negates the condition defined by the passed predicate.
func Not(predicate Predicate) Predicate {
	return PredicateFunc(func(v interface{}) bool {
		return !predicate.Accept(v)
	})
}

// TruePredicate is a predicate that returns true for all inputs.
func True() Predicate {
	return PredicateFunc(func(v interface{}) bool { return true })
}

// FalsePredicate is a predicate that returns false for all inputs.
func False() Predicate {
	return PredicateFunc(func(v interface{}) bool { return false })
}

// ExtractedValueAccepted returns A predicate that extracts a value using the Extractor and passes that value to the
// provided predicate
func ExtractedValueAccepted(extractor Extractor, predicate Predicate) Predicate {
	return PredicateFunc(func(v interface{}) bool {
		return predicate.Accept(extractor.Extract(v))
	})
}

// MethodIs returns a predicate that takes a request, extracts the method, and returns true if it equals the method
// provided, ignoring case.
func MethodIs(method string) Predicate {
	return ExtractedValueAccepted(UpperCaseExtractor(ExtractMethod()), StringEquals(strings.ToUpper(method)))
}
