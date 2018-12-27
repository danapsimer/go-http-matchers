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
	"regexp"
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

// StringEquals returns a predicate that returns true if the value passed is a string and is equal to the value of
// 'value'
func StringEquals(value string) Predicate {
	return PredicateFunc(func(s interface{}) bool {
		return s.(string) == value
	})
}

// StringContains returns a predicate that returns true if the value passed contains a substring matching 'value'.
func StringContains(value string) Predicate {
	return PredicateFunc(func(s interface{}) bool {
		return strings.Contains(s.(string), value)
	})
}

// StringStartsWith returns a predicate that returns true if the value passed starts with a substring matching 'value'.
func StringStartsWith(value string) Predicate {
	return PredicateFunc(func(s interface{}) bool {
		return strings.HasPrefix(s.(string), value)
	})
}

// StringEndsWith returns a predicate that returns true if the value passed ends with a substring matching 'value'.
func StringEndsWith(value string) Predicate {
	return PredicateFunc(func(s interface{}) bool {
		return strings.HasSuffix(s.(string), value)
	})
}

// StringMatches returns a predicate that returns true if the regex matches 'value'.
func StringMatches(regex *regexp.Regexp) Predicate {
	return PredicateFunc(func(s interface{}) bool {
		return regex.MatchString(s.(string))
	})
}

// ExtractedValueAccepted returns A predicate that extracts a value using the Extractor and passes that value to the
// provided predicate
func ExtractedValueAccepted(extractor Extractor, predicate Predicate) Predicate {
	return PredicateFunc(func(v interface{}) bool {
		return predicate.Accept(extractor.Extract(v))
	})
}

// PathMatches returns a predicate that returns true if the path matches the pathRegex.
func PathMatches(pathRegex *regexp.Regexp) Predicate {
	return ExtractedValueAccepted(ExtractPath(), StringMatches(pathRegex))
}

// PathEquals returns a predicate that returns true if the path equals 'path'
func PathEquals(path string) Predicate {
	return ExtractedValueAccepted(ExtractPath(), StringEquals(path))
}

// PathStartsWith returns a predicate that returns true if the path starts with 'path'
func PathStartsWith(path string) Predicate {
	return ExtractedValueAccepted(ExtractPath(), StringStartsWith(path))
}

// HeaderMatches returns a predicate that returns true if the header named 'name' matches 'regex'
func HeaderMatches(name string, regex *regexp.Regexp) Predicate {
	return ExtractedValueAccepted(ExtractHeader(name), StringMatches(regex))
}

// HeaderEquals returns a predicate that returns true if the header named 'name' equals 'value'
func HeaderEquals(name string, value string) Predicate {
	return ExtractedValueAccepted(ExtractHeader(name), StringEquals(value))
}

// HeaderEqualsIgnoreCase returns a predicate that returns true if the header named 'name' equals 'value', ignoring
// case.
func HeaderEqualsIgnoreCase(name string, path string) Predicate {
	return ExtractedValueAccepted(UpperCaseExtractor(ExtractHeader(name)), StringEquals(strings.ToUpper(path)))
}

// HeaderContains returns a predicate that returns true if the header named 'name' contains 'value'.
func HeaderContains(name string, path string) Predicate {
	return ExtractedValueAccepted(ExtractHeader(name), StringContains(path))
}

// HeaderContainsIgnoreCase returns a predicate that returns true if the header named 'name' contains 'value', ignoring
// case.
func HeaderContainsIgnoreCase(name string, path string) Predicate {
	return ExtractedValueAccepted(UpperCaseExtractor(ExtractHeader(name)), StringContains(strings.ToUpper(path)))
}

// HeaderStartsWith returns a predicate that returns true if the header named 'name' starts with 'value'.
func HeaderStartsWith(name string, path string) Predicate {
	return ExtractedValueAccepted(ExtractHeader(name), StringStartsWith(path))
}

// RequestURIMatches returns a predicate that returns true if the request URI matches the pathRegex.
func RequestURIMatches(pathRegex *regexp.Regexp) Predicate {
	return ExtractedValueAccepted(ExtractRequestURI(), StringMatches(pathRegex))
}

// RequestURIEquals returns a predicate that returns true if the request URI equals the path.
func RequestURIEquals(path string) Predicate {
	return ExtractedValueAccepted(ExtractRequestURI(), StringEquals(path))
}

// RequestURIStartsWith returns a predicate that returns true if the request URI starts with the path.
func RequestURIStartsWith(path string) Predicate {
	return ExtractedValueAccepted(ExtractRequestURI(), StringStartsWith(path))
}

// MethodIs returns a predicate that takes a request, extracts the method, and returns true if it equals the method
// provided, ignoring case.
func MethodIs(method string) Predicate {
	return ExtractedValueAccepted(UpperCaseExtractor(ExtractMethod()), StringEquals(strings.ToUpper(method)))
}

// QueryParamContainsIgnoreCase returns a Predicate that takes a request, extracts the query parameter specified and
// returns true if it equals the value provided.
func QueryParamEquals(name, value string) Predicate {
	return ExtractedValueAccepted(ExtractQueryParameter(name), StringEquals(value))
}

// QueryParamContainsIgnoreCase returns a Predicate that takes a request, extracts the query parameter specified and
// returns true if it equals the value provided, ignoring case.
func QueryParamEqualsIgnoreCase(name, value string) Predicate {
	return ExtractedValueAccepted(UpperCaseExtractor(ExtractQueryParameter(name)), StringEquals(strings.ToUpper(value)))
}

// QueryParamContainsIgnoreCase returns a Predicate that takes a request, extracts the query parameter specified and
// returns true if it contains the value provided.
func QueryParamContains(name, value string) Predicate {
	return ExtractedValueAccepted(ExtractQueryParameter(name), StringContains(value))
}

// QueryParamContainsIgnoreCase returns a Predicate that takes a request, extracts the query parameter specified and
// returns true if it contains the value provided, ignoring case.
func QueryParamContainsIgnoreCase(name, value string) Predicate {
	return ExtractedValueAccepted(UpperCaseExtractor(ExtractQueryParameter(name)), StringContains(strings.ToUpper(value)))
}

// QueryParamMatches returns a Predicate that takes a request, extracts the query parameter specified and
// returns true if the value matches the pattern provided.
func QueryParamMatches(name string, pattern *regexp.Regexp) Predicate {
	return ExtractedValueAccepted(ExtractQueryParameter(name), StringMatches(pattern))
}

// QueryParamStartsWith returns a Predicate that takes a request, extracts the query parameter specified and
// returns true if the value starts with the prefix provided.
func QueryParamStartsWith(name string, prefix string) Predicate {
	return ExtractedValueAccepted(ExtractQueryParameter(name), StringStartsWith(prefix))
}
