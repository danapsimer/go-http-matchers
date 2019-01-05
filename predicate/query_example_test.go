package predicate_test

import (
	"fmt"
	. "github.com/bluesoftdev/go-http-matchers/predicate"
	"net/http"
	"regexp"
)

func ExampleQueryParamContains() {
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=foobar&l=3", nil)
	fmt.Printf("%v\n", QueryParamContains("q", "oob").Accept(req))
	fmt.Printf("%v\n", QueryParamContains("q", "snafu").Accept(req))
	fmt.Printf("%v\n", QueryParamContains("x", "oob").Accept(req))
	// Output:
	// true
	// false
	// false
}

func ExampleQueryParamContainsIgnoreCase() {
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=foobar&l=3", nil)
	fmt.Printf("%v\n", QueryParamContainsIgnoreCase("q", "OoB").Accept(req))
	fmt.Printf("%v\n", QueryParamContainsIgnoreCase("q", "snafu").Accept(req))
	fmt.Printf("%v\n", QueryParamContainsIgnoreCase("x", "OoB").Accept(req))
	// Output:
	// true
	// false
	// false
}

func ExampleQueryParamEquals() {
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=foobar&l=3", nil)
	fmt.Printf("%v\n", QueryParamEquals("q", "foobar").Accept(req))
	fmt.Printf("%v\n", QueryParamEquals("q", "snafu").Accept(req))
	fmt.Printf("%v\n", QueryParamEquals("x", "foobar").Accept(req))
	// Output:
	// true
	// false
	// false
}

func ExampleQueryParamEqualsIgnoreCase() {
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=foobar&l=3", nil)
	fmt.Printf("%v\n", QueryParamEqualsIgnoreCase("q", "FooBar").Accept(req))
	fmt.Printf("%v\n", QueryParamEqualsIgnoreCase("q", "Snafu").Accept(req))
	fmt.Printf("%v\n", QueryParamEqualsIgnoreCase("x", "FooBar").Accept(req))
	// Output:
	// true
	// false
	// false
}

func ExampleQueryParamStartsWith() {
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=foobar&l=3", nil)
	fmt.Printf("%v\n", QueryParamStartsWith("q", "foo").Accept(req))
	fmt.Printf("%v\n", QueryParamStartsWith("q", "snafu").Accept(req))
	fmt.Printf("%v\n", QueryParamStartsWith("x", "foo").Accept(req))
	// Output:
	// true
	// false
	// false
}

func ExampleQueryParamMatches() {
	truePattern := regexp.MustCompile("fo{2}bar")
	falsePattern := regexp.MustCompile("barfo{2}")
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=foobar&l=3", nil)
	fmt.Printf("%v\n", QueryParamMatches("q", truePattern).Accept(req))
	fmt.Printf("%v\n", QueryParamMatches("q", falsePattern).Accept(req))
	// Output:
	// true
	// false
}
