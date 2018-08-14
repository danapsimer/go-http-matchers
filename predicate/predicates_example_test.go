package predicate_test
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
	"fmt"
	"gitlab.com/ComputersFearMe/go-http-matchers/predicate"
	"net/http"
	"regexp"
)

func ExampleNot() {
	fmt.Printf("%v\n",predicate.Not(predicate.True()).Accept(nil))
	fmt.Printf("%v\n",predicate.Not(predicate.False()).Accept(nil))
	// Output:
	// false
	// true
}

func ExampleAnd() {
	fmt.Printf("%v\n",predicate.And(predicate.True(),predicate.True()).Accept(nil))
	fmt.Printf("%v\n",predicate.And(predicate.False(),predicate.True()).Accept(nil))
	// Output:
	// true
	// false
}

func ExampleOr() {
	fmt.Printf("%v\n",predicate.Or(predicate.False(),predicate.False()).Accept(nil))
	fmt.Printf("%v\n",predicate.Or(predicate.False(),predicate.True()).Accept(nil))
	// Output:
	// false
	// true
}

func ExampleRequestURIStartsWith() {
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	fmt.Printf("%v\n",predicate.RequestURIStartsWith("/test/foo").Accept(req))
	fmt.Printf("%v\n",predicate.RequestURIStartsWith("/test/bar").Accept(req))
	// Output:
	// true
	// false
}

func ExampleRequestURIEquals() {
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	fmt.Printf("%v\n",predicate.RequestURIEquals("/test/foo/bar?q=5&l=3").Accept(req))
	fmt.Printf("%v\n",predicate.RequestURIEquals("/test/foo/bar?q=6&l=3").Accept(req))
	// Output:
	// true
	// false
}

func ExampleRequestURIMatches() {
	truePattern := regexp.MustCompile("bar\\?q=\\d*")
	falsePattern := regexp.MustCompile("foo\\?q\\d*")
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	fmt.Printf("%v\n",predicate.RequestURIMatches(truePattern).Accept(req))
	fmt.Printf("%v\n",predicate.RequestURIMatches(falsePattern).Accept(req))
	// Output:
	// true
	// false
}
