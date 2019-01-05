package predicate_test

import (
	"fmt"
	. "github.com/bluesoftdev/go-http-matchers/predicate"
	"regexp"
)

func ExampleStringContains() {
	fmt.Printf("%v\n", StringContains("oob").Accept("foobar"))
	fmt.Printf("%v\n", StringContains("snafu").Accept("foobar"))
	// Output:
	// true
	// false
}

func ExampleStringEndsWith() {
	fmt.Printf("%v\n", StringEndsWith("bar").Accept("foobar"))
	fmt.Printf("%v\n", StringEndsWith("foo").Accept("foobar"))
	// Output:
	// true
	// false
}

func ExampleStringStartsWith() {
	fmt.Printf("%v\n", StringStartsWith("foo").Accept("foobar"))
	fmt.Printf("%v\n", StringStartsWith("bar").Accept("foobar"))
	// Output:
	// true
	// false
}

func ExampleStringEquals() {
	fmt.Printf("%v\n", StringEquals("foobar").Accept("foobar"))
	fmt.Printf("%v\n", StringEquals("barfoo").Accept("foobar"))
	// Output:
	// true
	// false
}

func ExampleStringMatches() {
	truePattern := regexp.MustCompile("fo{2}bar")
	falsePattern := regexp.MustCompile("barfo{2}")
	fmt.Printf("%v\n", StringMatches(truePattern).Accept("foobar"))
	fmt.Printf("%v\n", StringMatches(falsePattern).Accept("foobar"))
	// Output:
	// true
	// false
}
