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
	"github.com/stretchr/testify/assert"
	"go-http-matchers/extractor"
	"net/http"
	"net/http/httputil"
	"reflect"
	"strings"
	"testing"
)

func TestExtractHeader(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	req.Header.Add("Foo", "Bar")

	result := extractor.ExtractHeader("Foo").Extract(req)
	assert.Equal(t, "Bar", result, "expected result to be Bar but got: "+result.(string))
}

func TestExtractHeader_NotPresent(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")

	result := extractor.ExtractHeader("Foo").Extract(req)
	assert.Equal(t, "", result, "expected result to be empty but got: "+result.(string))
}

func TestExtractQueryParameter_Q(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")

	result := extractor.ExtractQueryParameter("q").Extract(req)
	assert.Equal(t, "5", result, "expected result to be empty but got: "+result.(string))
}

func TestExtractQueryParameter_L(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")

	result := extractor.ExtractQueryParameter("l").Extract(req)
	assert.Equal(t, "3", result, "expected result to be empty but got: "+result.(string))
}

func TestExtractQueryParameter_NotPresent(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")

	result := extractor.ExtractQueryParameter("z").Extract(req)
	assert.Equal(t, "", result, "expected result to be empty but got: "+result.(string))
}

func TestUpperCaseExtractor(t *testing.T) {
	result := extractor.UpperCaseExtractor(extractor.ExtractorFunc(func(interface{}) interface{} {
		return "foo"
	})).Extract(nil)

	assert.Equal(t, "FOO", result, "Expected result to be upper cased.")
}

func TestUpperCaseExtractor_ReturnsNil(t *testing.T) {
	result := extractor.UpperCaseExtractor(extractor.ExtractorFunc(func(interface{}) interface{} {
		return nil
	})).Extract(nil)

	assert.Nil(t, result, "Expected result to be upper cased.")
}

func TestExtractPathElementByIndex_atBegining(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")

	result := extractor.ExtractPathElementByIndex(1).Extract(req)
	assert.Equal(t, "test", result, "expected result to be 'test' but got: "+result.(string))
}

func TestExtractPathElementByIndex_inMiddle(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")

	result := extractor.ExtractPathElementByIndex(2).Extract(req)
	assert.Equal(t, "foo", result, "expected result to be 'foo' but got: "+result.(string))
}

func TestExtractPathElementByIndex_atEnd(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")

	result := extractor.ExtractPathElementByIndex(3).Extract(req)
	assert.Equal(t, "bar", result, "expected result to be 'bar' but got: "+result.(string))
}
func TestExtractPathElementByIndex_pastEnd(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")

	result := extractor.ExtractPathElementByIndex(4).Extract(req)
	assert.Equal(t, "", result, "expected result to be empty but got: "+result.(string))
}

func TestExtractPathElementByIndex_FromEnd(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")

	result := extractor.ExtractPathElementByIndex(-1).Extract(req)
	assert.Equal(t, "bar", result, "expected result to be 'bar' but got: "+result.(string))
}

func TestExtractPathElementByIndex_FromEndPastBeginning(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")

	result := extractor.ExtractPathElementByIndex(-5).Extract(req)
	assert.Equal(t, "", result, "expected result to be empty but got: "+result.(string))
}

const testXml = `
<foo>
  <bar snafu="foobar"/>
</foo>
`

func TestExtractXPathString(t *testing.T) {
	req, err := http.NewRequest("POST", "http://foo.com/test", strings.NewReader(testXml))
	assert.NoError(t, err, "failed to create test request.")

	result := extractor.ExtractXPathString("/foo/bar/@snafu").Extract(req)
	assert.Equal(t, "foobar", result, "expected result to be 'foobar' but got: "+result.(string))
}

func TestExtractXPathString_NotPresent(t *testing.T) {
	req, err := http.NewRequest("POST", "http://foo.com/test", strings.NewReader(testXml))
	assert.NoError(t, err, "failed to create test request.")

	result := extractor.ExtractXPathString("/foo/bar/@fubar").Extract(req)
	assert.Equal(t, "", result, "expected result to be 'foobar' but got: "+result.(string))
}

func TestExtractRequestURI(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")

	result := extractor.ExtractRequestURI().Extract(req)
	assert.Equal(t, "/test/foo/bar?q=5&l=3", result, "expected result to be '/test/foo/bar?q=5&l=3' but got: "+result.(string))
}

func TestExtractPath(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")

	result := extractor.ExtractPath().Extract(req)
	assert.Equal(t, "/test/foo/bar", result, "expected result to be '/test/foo/bar' but got: "+result.(string))
}

func TestExtractMethod(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")

	result := extractor.ExtractMethod().Extract(req)
	assert.Equal(t, "GET", result, "expected result to be 'GET' but got: "+result.(string))
}

func dumpReq(req *http.Request) string {
	dump, err := httputil.DumpRequest(req, false)
	if err != nil {
		return "ERROR dumping request: " + err.Error()
	}
	return fmt.Sprintf("%q", dump)
}

func TestExtractIdentity(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")

	result := extractor.IdentityExtractor().Extract(req)
	if (assert.IsType(t, &http.Request{}, result, "Expected an *http.Request but got: "+reflect.TypeOf(result).String())) {
		resultReq := result.(*http.Request)
		assert.Equal(t, req, resultReq,
			fmt.Sprintf("expected result to be the same as the passed in request but got: %s", dumpReq(resultReq)))
	}
}

func TestExtractHost(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	result := extractor.ExtractHost().Extract(req)
	assert.Equal(t, "foo.com", result)
}
