package predicate_test

import (
	. "github.com/bluesoftdev/go-http-matchers/predicate"
	"github.com/stretchr/testify/assert"
	"net/http"
	"regexp"
	"testing"
)

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

