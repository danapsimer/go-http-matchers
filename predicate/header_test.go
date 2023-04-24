package predicate_test

import (
	. "github.com/bluesoftdev/go-http-matchers/predicate"
	"github.com/stretchr/testify/assert"
	"net/http"
	"regexp"
	"testing"
)

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

func TestRequestKeyStringMatches(t *testing.T) {
	key := "foo"
	assert.False(t, StringMatches(regexp.MustCompile("\\d+")).Accept(key))
	assert.True(t, StringMatches(regexp.MustCompile("[a-z]+")).Accept(key))
}
