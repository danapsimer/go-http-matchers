package predicate_test

import (
	"fmt"
	"github.com/danapsimer/go-http-matchers/predicate"
	"net/http"
	"regexp"
)

func ExampleRequestURIStartsWith() {
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	fmt.Printf("%v\n", predicate.RequestURIStartsWith("/test/foo").Accept(req))
	fmt.Printf("%v\n", predicate.RequestURIStartsWith("/test/bar").Accept(req))
	// Output:
	// true
	// false
}

func ExampleRequestURIEquals() {
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	fmt.Printf("%v\n", predicate.RequestURIEquals("/test/foo/bar?q=5&l=3").Accept(req))
	fmt.Printf("%v\n", predicate.RequestURIEquals("/test/foo/bar?q=6&l=3").Accept(req))
	// Output:
	// true
	// false
}

func ExampleRequestURIMatches() {
	truePattern := regexp.MustCompile("bar\\?q=\\d*")
	falsePattern := regexp.MustCompile("foo\\?q\\d*")
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	fmt.Printf("%v\n", predicate.RequestURIMatches(truePattern).Accept(req))
	fmt.Printf("%v\n", predicate.RequestURIMatches(falsePattern).Accept(req))
	// Output:
	// true
	// false
}
