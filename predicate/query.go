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
	. "go-http-matchers/extractor"
	"regexp"
	"strings"
)

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
