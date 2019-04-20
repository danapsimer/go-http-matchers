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
)

// RequestURIMatches returns a predicate that returns true if the request URI matches the pathRegex.
func RequestURIMatches(pathRegex *regexp.Regexp) Predicate {
	return ExtractedValueAccepted(extractor.ExtractRequestURI(), StringMatches(pathRegex))
}

// RequestURIEquals returns a predicate that returns true if the request URI equals the path.
func RequestURIEquals(path string) Predicate {
	return ExtractedValueAccepted(extractor.ExtractRequestURI(), StringEquals(path))
}

// RequestURIStartsWith returns a predicate that returns true if the request URI starts with the path.
func RequestURIStartsWith(path string) Predicate {
	return ExtractedValueAccepted(extractor.ExtractRequestURI(), StringStartsWith(path))
}
