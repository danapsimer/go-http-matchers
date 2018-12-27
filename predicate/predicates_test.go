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
	. "go-http-matchers/predicate"
	"github.com/stretchr/testify/assert"
	"net/http"
	"regexp"
	"testing"
)

func TestTrue(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	assert.True(t, True().Accept(req), "expected true.")
}

func TestFalse(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	assert.False(t, False().Accept(req), "expected false.")
}

func TestNot(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	assert.False(t, Not(True()).Accept(req))
	assert.True(t, Not(False()).Accept(req))
}

func TestAnd(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	assert.True(t, And(True(), True()).Accept(req), "expected true.")
	assert.True(t, And(True(), True(), True()).Accept(req), "expected false.")
	assert.False(t, And(False(), True()).Accept(req), "expected false.")
	assert.False(t, And(True(), False()).Accept(req), "expected false.")
	assert.False(t, And(True(), True(), False()).Accept(req), "expected false.")
	assert.False(t, And(True(), False(), True()).Accept(req), "expected false.")
	assert.False(t, And(False(), True(), True()).Accept(req), "expected false.")
}

func TestOr(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	assert.True(t, Or(True(), True()).Accept(req), "expected true.")
	assert.True(t, Or(True(), True(), True()).Accept(req), "expected true.")
	assert.True(t, Or(False(), True()).Accept(req), "expected true.")
	assert.True(t, Or(True(), False()).Accept(req), "expected true.")
	assert.True(t, Or(True(), True(), False()).Accept(req), "expected true.")
	assert.True(t, Or(True(), False(), True()).Accept(req), "expected true.")
	assert.True(t, Or(False(), True(), True()).Accept(req), "expected true.")
	assert.False(t, Or(False(), False()).Accept(req), "expected false.")
	assert.False(t, Or(False(), False(), False()).Accept(req), "expected false.")
}

func TestHeaderContains(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	req.Header.Add("FOO", "FOOBAR")

	assert.True(t, HeaderContainsIgnoreCase("FOO", "BAR").Accept(req), "expected true.")
	assert.False(t, HeaderContainsIgnoreCase("FOO", "snafu").Accept(req), "expected false.")
}

func TestHeaderContainsIgnoreCase(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	req.Header.Add("FOO", "FOOBAR")

	assert.True(t, HeaderContainsIgnoreCase("FOO", "bar").Accept(req), "expected true.")
	assert.False(t, HeaderContainsIgnoreCase("FOO", "snafu").Accept(req), "expected false.")
}

func TestHeaderEquals(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	req.Header.Add("FOO", "FOOBAR")

	assert.True(t, HeaderEquals("FOO", "FOOBAR").Accept(req), "expected true.")
	assert.False(t, HeaderEquals("FOO", "snafu").Accept(req), "expected false.")
}

func TestHeaderEqualsIgnoreCase(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	req.Header.Add("FOO", "FOOBAR")

	assert.True(t, HeaderEqualsIgnoreCase("FOO", "FooBar").Accept(req), "expected true.")
	assert.False(t, HeaderEqualsIgnoreCase("FOO", "snafu").Accept(req), "expected false.")
}

func TestHeaderMatches(t *testing.T) {
	truePattern := regexp.MustCompile("FO{2}B.*")
	falsePattern := regexp.MustCompile("FO{3}B.*")
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	req.Header.Add("FOO", "FOOBAR")

	assert.True(t, HeaderMatches("FOO", truePattern).Accept(req), "expected true.")
	assert.False(t, HeaderMatches("FOO", falsePattern).Accept(req), "expected false.")
}

func TestHeaderStartsWith(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	req.Header.Add("FOO", "FOOBAR")

	assert.True(t, HeaderStartsWith("FOO", "FOO").Accept(req), "expected true.")
	assert.False(t, HeaderStartsWith("FOO", "snafu").Accept(req), "expected false.")
}

func TestRequestURIMatches(t *testing.T) {
	truePattern := regexp.MustCompile("bar\\?q=5")
	falsePattern := regexp.MustCompile("bar\\?x=5")
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")

	assert.True(t, RequestURIMatches(truePattern).Accept(req), "expected true.")
	assert.False(t, RequestURIMatches(falsePattern).Accept(req), "expected false.")
}

func TestRequestURIStartsWith(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	assert.True(t, RequestURIStartsWith("/test/foo").Accept(req))
	assert.False(t, RequestURIStartsWith("/test/bar").Accept(req))
}

func TestRequestURIEquals(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	assert.True(t, RequestURIEquals("/test/foo/bar?q=5&l=3").Accept(req))
	assert.False(t, RequestURIEquals("/test/bar/foo?q=5&l=3").Accept(req))
}

func TestMethodIs(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	assert.True(t, MethodIs("GET").Accept(req))
	assert.False(t, MethodIs("POST").Accept(req))
}

func TestPathEquals(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	assert.True(t, PathEquals("/test/foo/bar").Accept(req))
	assert.False(t, PathEquals("/test/bar/foo").Accept(req))
}

func TestPathMatches(t *testing.T) {
	truePattern := regexp.MustCompile("fo{2}/bar")
	falsePattern := regexp.MustCompile("bar/fo{2}")
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	assert.True(t, PathMatches(truePattern).Accept(req))
	assert.False(t, PathMatches(falsePattern).Accept(req))
}

func TestPathStartsWith(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	assert.True(t, PathStartsWith("/test/foo/").Accept(req))
	assert.False(t, PathStartsWith("/test/bar/").Accept(req))
}

func TestQueryParamContains(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=foobar&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	assert.True(t, QueryParamContains("q", "oob").Accept(req))
	assert.False(t, QueryParamContains("q", "snafu").Accept(req))
	assert.False(t, QueryParamContains("x", "oob").Accept(req))
}

func TestQueryParamContainsIgnoreCase(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=foobar&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	assert.True(t, QueryParamContainsIgnoreCase("q", "OoB").Accept(req))
	assert.False(t, QueryParamContainsIgnoreCase("q", "snafu").Accept(req))
	assert.False(t, QueryParamContainsIgnoreCase("x", "OoB").Accept(req))
}

func TestQueryParamEquals(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=foobar&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	assert.True(t, QueryParamEquals("q", "foobar").Accept(req))
	assert.False(t, QueryParamEquals("q", "snafu").Accept(req))
	assert.False(t, QueryParamEquals("x", "foobar").Accept(req))
}

func TestQueryParamEqualsIgnoreCase(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=foobar&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	assert.True(t, QueryParamEqualsIgnoreCase("q", "FooBar").Accept(req))
	assert.False(t, QueryParamEqualsIgnoreCase("q", "Snafu").Accept(req))
	assert.False(t, QueryParamEqualsIgnoreCase("x", "FooBar").Accept(req))
}

func TestQueryParamStartsWith(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=foobar&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	assert.True(t, QueryParamStartsWith("q", "foo").Accept(req))
	assert.False(t, QueryParamStartsWith("q", "snafu").Accept(req))
	assert.False(t, QueryParamStartsWith("x", "foo").Accept(req))
}

func TestQueryParamMatches(t *testing.T) {
	truePattern := regexp.MustCompile("fo{2}bar")
	falsePattern := regexp.MustCompile("barfo{2}")
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=foobar&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	assert.True(t, QueryParamMatches("q", truePattern).Accept(req))
	assert.False(t, QueryParamMatches("q", falsePattern).Accept(req))
}

func TestStringContains(t *testing.T) {
	assert.True(t, StringContains("oob").Accept("foobar"))
	assert.False(t, StringContains("snafu").Accept("foobar"))
}

func TestStringEndsWith(t *testing.T) {
	assert.True(t, StringEndsWith("bar").Accept("foobar"))
	assert.False(t, StringEndsWith("foo").Accept("foobar"))
}

func TestStringStartsWith(t *testing.T) {
	assert.True(t, StringStartsWith("foo").Accept("foobar"))
	assert.False(t, StringStartsWith("bar").Accept("foobar"))
}

func TestStringEquals(t *testing.T) {
	assert.True(t, StringEquals("foobar").Accept("foobar"))
	assert.False(t, StringEquals("barfoo").Accept("foobar"))
}

func TestStringMatches(t *testing.T) {
	truePattern := regexp.MustCompile("fo{2}bar")
	falsePattern := regexp.MustCompile("barfo{2}")
	assert.True(t, StringMatches(truePattern).Accept("foobar"))
	assert.False(t, StringMatches(falsePattern).Accept("foobar"))
}
