package predicate_test

import (
	"fmt"
	"github.com/danapsimer/go-http-matchers/predicate"
	"net/http"
	"regexp"
)

func ExampleHeaderContains() {
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	req.Header.Add("FOO", "FOOBAR")

	fmt.Printf("%v\n", predicate.HeaderContains("FOO", "BAR").Accept(req))
	fmt.Printf("%v\n", predicate.HeaderContains("FOO", "snafu").Accept(req))
	// Output:
	// true
	// false
}

func ExampleHeaderContainsIgnoreCase() {
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	req.Header.Add("FOO", "FOOBAR")

	fmt.Printf("%v\n", predicate.HeaderContainsIgnoreCase("FOO", "bar").Accept(req))
	fmt.Printf("%v\n", predicate.HeaderContainsIgnoreCase("FOO", "snafu").Accept(req))
	// Output:
	// true
	// false
}

func ExampleHeaderEquals() {
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	req.Header.Add("FOO", "FOOBAR")

	fmt.Printf("%v\n", predicate.HeaderEquals("FOO", "FOOBAR").Accept(req))
	fmt.Printf("%v\n", predicate.HeaderEquals("FOO", "snafu").Accept(req))
	// Output:
	// true
	// false
}

func ExampleHeaderEqualsIgnoreCase() {
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	req.Header.Add("FOO", "FOOBAR")

	fmt.Printf("%v\n", predicate.HeaderEqualsIgnoreCase("FOO", "FooBar").Accept(req))
	fmt.Printf("%v\n", predicate.HeaderEqualsIgnoreCase("FOO", "snafu").Accept(req))
	// Output:
	// true
	// false
}

func ExampleHeaderMatches() {
	truePattern := regexp.MustCompile("FO{2}B.*")
	falsePattern := regexp.MustCompile("FO{3}B.*")
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	req.Header.Add("FOO", "FOOBAR")

	fmt.Printf("%v\n", predicate.HeaderMatches("FOO", truePattern).Accept(req))
	fmt.Printf("%v\n", predicate.HeaderMatches("FOO", falsePattern).Accept(req))
	// Output:
	// true
	// false
}

func ExampleHeaderStartsWith() {
	req, _ := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	req.Header.Add("FOO", "FOOBAR")

	fmt.Printf("%v\n", predicate.HeaderStartsWith("FOO", "FOO").Accept(req))
	fmt.Printf("%v\n", predicate.HeaderStartsWith("FOO", "snafu").Accept(req))
	// Output:
	// true
	// false
}
