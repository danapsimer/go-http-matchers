package predicate_test

import (
	. "github.com/danapsimer/go-http-matchers/predicate"
	"github.com/stretchr/testify/assert"
	"net/http"
	"regexp"
	"testing"
)

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
