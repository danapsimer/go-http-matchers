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
)

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
