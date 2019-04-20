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
	"regexp"
	"strings"
)

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
