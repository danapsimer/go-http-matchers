// Package extractor contains higher order functions that return functions that can take a value and return some
// derivative value for it.  It also defines a number of convenient extractors for use with http.Request objects.  The
// idea is that the Extractor producing functions can take inputs that affect the derivative function's output.  For
// instance the ExtractHeader(name string) function returns an instance of Extractor that expects a *http.Request to be
// passed and returns the Header with the name passed to the higher order function.  In this way, the extractor can
// be called repeatedly with different requests.  In most cases, Extractors are used with Predicates.
package extractor
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
	"gopkg.in/xmlpath.v2"
	"net/http"
	"strings"
)

// Extractor can extract a value from another value by calling the Extract method.
type Extractor interface {
	Extract(interface{}) interface{}
}

// ExtractorFunc is a function that calls itself when it's Extract method is called.
type ExtractorFunc func(interface{}) interface{}

func (ef ExtractorFunc) Extract(v interface{}) interface{} {
	return ef(v)
}

// IdentityExtractor returns an Extractor that returns the value passed.
func IdentityExtractor() Extractor {
	return ExtractorFunc(func(r interface{}) interface{} {
		return r
	})
}

// ExtractMethod returns an extractor that expects a *http.Request and returns the method.
func ExtractMethod() Extractor {
	return ExtractorFunc(func(r interface{}) interface{} {
		return r.(*http.Request).Method
	})
}

// ExtractPath returns an Extractor that expects a *http.Request and returns the URL's Path property.
func ExtractPath() Extractor {
	return ExtractorFunc(func(r interface{}) interface{} {
		return r.(*http.Request).URL.Path
	})
}

// ExtractRequestURI returns an Extractor that expects a *http.Request and returns the URL's RequestURI property.
func ExtractRequestURI() Extractor {
	return ExtractorFunc(func(r interface{}) interface{} {
		return r.(*http.Request).URL.RequestURI()
	})
}

// ExtractHeader returns an Extractor that expects a *http.Request and returns the value of the header named 'name'.
func ExtractHeader(name string) Extractor {
	return ExtractorFunc(func(r interface{}) interface{} {
		if "HOST" == strings.ToUpper(name) {
			return r.(*http.Request).Host
		}
		return r.(*http.Request).Header.Get(name)
	})
}

// ExtractHost returns an Extractor that returns the value of the "Host" element in the request.
func ExtractHost() Extractor {
	return ExtractorFunc(func(r interface{}) interface{} {
		return r.(*http.Request).Host
	})
}

// UpperCaseExtractor returns an Extractor that decorates the passed extractor by applying strings.ToUpper to the
// value returned.
func UpperCaseExtractor(extractor Extractor) Extractor {
	return ExtractorFunc(func(v interface{}) interface{} {
		value := extractor.Extract(v)
		if value == nil {
			return nil
		}
		return strings.ToUpper(value.(string))
	})
}

// ExtractXPathString returns a Extractor that expects a *http.Request and uses the passed XPath expression to extract
// a string from the Body of the request Request.
func ExtractXPathString(xpath string) Extractor {
	path := xmlpath.MustCompile(xpath)
	return ExtractorFunc(func(r interface{}) interface{} {
		str := ""
		root, err := xmlpath.Parse(r.(*http.Request).Body)
		if err == nil {
			str, _ = path.String(root)
		}
		return str
	})
}

// ExtractPathElementByIndex returns an Extractor that expects a *http.Request and extracts the path element at the
// given position.  A negative number denotes a position from the end (starting at 1 e.g. -1 is the last element in the
// path). For positive inputs, the counting starts at 1 as well.
func ExtractPathElementByIndex(idx int) Extractor {
	return ExtractorFunc(func(r interface{}) interface{} {
		elements := strings.Split(r.(*http.Request).URL.Path, "/")
		var i int
		if idx < 0 {
			i = len(elements) + idx
		} else {
			i = idx
		}
		if i < 0 || i >= len(elements) {
			return ""
		}
		return elements[i]
	})
}

// ExtractQueryParameter returns an Extractor that expects a *http.Request and extracts they named query parameter's
// value.
func ExtractQueryParameter(name string) Extractor {
	return ExtractorFunc(func(r interface{}) interface{} {
		return r.(*http.Request).URL.Query().Get(name)
	})
}
