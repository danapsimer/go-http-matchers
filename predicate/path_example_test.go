package predicate_test

import (
	"fmt"
	. "github.com/bluesoftdev/go-http-matchers/predicate"
	"net/http"
	"regexp"
)

func ExamplePathEquals() {
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	fmt.Printf("%v\n", PathEquals("/test/foo/bar").Accept(req))
	fmt.Printf("%v\n", PathEquals("/test/bar/foo").Accept(req))
	// Output:
	// true
	// false
}

func ExamplePathMatches() {
	truePattern := regexp.MustCompile("fo{2}/bar")
	falsePattern := regexp.MustCompile("bar/fo{2}")
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	fmt.Printf("%v\n", PathMatches(truePattern).Accept(req))
	fmt.Printf("%v\n", PathMatches(falsePattern).Accept(req))
	// Output:
	// true
	// false
}

func ExamplePathStartsWith() {
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	fmt.Printf("%v\n", PathStartsWith("/test/foo/").Accept(req))
	fmt.Printf("%v\n", PathStartsWith("/test/bar/").Accept(req))
	// Output:
	// true
	// false
}
