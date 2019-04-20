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
	"github.com/bluesoftdev/go-http-matchers/extractor"
	"regexp"
	"strings"
)

// HeaderMatches returns a predicate that returns true if the header named 'name' matches 'regex'
func HeaderMatches(name string, regex *regexp.Regexp) Predicate {
	return ExtractedValueAccepted(extractor.ExtractHeader(name), StringMatches(regex))
}

// HeaderEquals returns a predicate that returns true if the header named 'name' equals 'value'
func HeaderEquals(name string, value string) Predicate {
	return ExtractedValueAccepted(extractor.ExtractHeader(name), StringEquals(value))
}

// HeaderEqualsIgnoreCase returns a predicate that returns true if the header named 'name' equals 'value', ignoring
// case.
func HeaderEqualsIgnoreCase(name string, path string) Predicate {
	return ExtractedValueAccepted(extractor.UpperCaseExtractor(extractor.ExtractHeader(name)), StringEquals(strings.ToUpper(path)))
}

// HeaderContains returns a predicate that returns true if the header named 'name' contains 'value'.
func HeaderContains(name string, path string) Predicate {
	return ExtractedValueAccepted(extractor.ExtractHeader(name), StringContains(path))
}

// HeaderContainsIgnoreCase returns a predicate that returns true if the header named 'name' contains 'value', ignoring
// case.
func HeaderContainsIgnoreCase(name string, path string) Predicate {
	return ExtractedValueAccepted(extractor.UpperCaseExtractor(extractor.ExtractHeader(name)), StringContains(strings.ToUpper(path)))
}

// HeaderStartsWith returns a predicate that returns true if the header named 'name' starts with 'value'.
func HeaderStartsWith(name string, path string) Predicate {
	return ExtractedValueAccepted(extractor.ExtractHeader(name), StringStartsWith(path))
}
