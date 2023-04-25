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
	"github.com/danapsimer/go-http-matchers/predicate"
)

func ExampleNot() {
	fmt.Printf("%v\n", predicate.Not(predicate.True()).Accept(nil))
	fmt.Printf("%v\n", predicate.Not(predicate.False()).Accept(nil))
	// Output:
	// false
	// true
}

func ExampleAnd() {
	fmt.Printf("%v\n", predicate.And(predicate.True(), predicate.True()).Accept(nil))
	fmt.Printf("%v\n", predicate.And(predicate.False(), predicate.True()).Accept(nil))
	// Output:
	// true
	// false
}

func ExampleOr() {
	fmt.Printf("%v\n", predicate.Or(predicate.False(), predicate.False()).Accept(nil))
	fmt.Printf("%v\n", predicate.Or(predicate.False(), predicate.True()).Accept(nil))
	// Output:
	// false
	// true
}
