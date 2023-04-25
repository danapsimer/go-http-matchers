package predicate_test

import (
	. "github.com/danapsimer/go-http-matchers/predicate"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

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
