package extractor_test

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
	"github.com/danapsimer/go-http-matchers/extractor"
	"net/http"
	"net/http/httputil"
	"strings"
)

func ExampleExtractPath() {
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	fmt.Printf("path = %s", extractor.ExtractPath().Extract(req))
	// Output: path = /test/foo/bar
}

func ExampleExtractMethod() {
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	fmt.Printf("method = %s", extractor.ExtractMethod().Extract(req))
	// Output: method = GET
}

func ExampleExtractRequestURI() {
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	fmt.Printf("requestURI = %s", extractor.ExtractRequestURI().Extract(req))
	// Output: requestURI = /test/foo/bar?q=5&l=3
}

func ExampleExtractHeader() {
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	req.Header.Add("FOO", "FOOBAR")
	fmt.Printf("header[FOO] = %s", extractor.ExtractHeader("FOO").Extract(req))
	// Output: header[FOO] = FOOBAR
}

func ExampleExtractPathElementByIndex() {
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	for i := 0; i <= 4; i++ {
		fmt.Printf("PathElementByIndex[%d] = '%s'\n", i, extractor.ExtractPathElementByIndex(i).Extract(req))
	}
	for i := -1; i >= -4; i-- {
		fmt.Printf("PathElementByIndex[%d] = '%s'\n", i, extractor.ExtractPathElementByIndex(i).Extract(req))
	}
	// Output:
	// PathElementByIndex[0] = ''
	// PathElementByIndex[1] = 'test'
	// PathElementByIndex[2] = 'foo'
	// PathElementByIndex[3] = 'bar'
	// PathElementByIndex[4] = ''
	// PathElementByIndex[-1] = 'bar'
	// PathElementByIndex[-2] = 'foo'
	// PathElementByIndex[-3] = 'test'
	// PathElementByIndex[-4] = ''
}

func ExampleIdentityExtractor() {
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	dump, _ := httputil.DumpRequest(extractor.IdentityExtractor().Extract(req).(*http.Request), false)
	fmt.Printf("request = %q", dump)
	// Output: request = "GET /test/foo/bar?q=5&l=3 HTTP/1.1\r\nHost: foo.com\r\n\r\n"
}

func ExampleExtractQueryParameter() {
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	fmt.Printf("query[q] = %s\n", extractor.ExtractQueryParameter("q").Extract(req))
	fmt.Printf("query[l] = %s\n", extractor.ExtractQueryParameter("l").Extract(req))
	fmt.Printf("query[x] = %s\n", extractor.ExtractQueryParameter("x").Extract(req))
	// Output:
	// query[q] = 5
	// query[l] = 3
	// query[x] =
}

func ExampleExtractXPathString() {
	const testXml = `
<foo>
  <bar snafu="foobar"/>
</foo>
`
	req, _ := http.NewRequest("POST", "http://foo.com/test", strings.NewReader(testXml))
	fmt.Printf("xpath[/foo/bar/@snafu] = %s", extractor.ExtractXPathString("/foo/bar/@snafu").Extract(req))
	// Output: xpath[/foo/bar/@snafu] = foobar
}

func ExampleUpperCaseExtractor() {
	fmt.Printf("upperCase[FooBar] = %s", extractor.UpperCaseExtractor(extractor.IdentityExtractor()).Extract("FooBar"))
	// Output: upperCase[FooBar] = FOOBAR
}
