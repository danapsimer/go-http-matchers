package predicate_test

import (
	. "github.com/danapsimer/go-http-matchers/predicate"
	"github.com/stretchr/testify/assert"
	"net/http"
	"regexp"
	"testing"
)

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
