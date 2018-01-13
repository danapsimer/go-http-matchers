package predicate

import (
	. "github.com/bluesoftdev/go-http-matchers/extractor"
	"strings"
	"regexp"
)

// Predicate is a class that can accept or reject a value based on some condition.
type Predicate interface {
	Accept(interface{}) bool
}

// PredicateFunc is an implementation of Predicate that is a function and calls itself on a call to Accept
type PredicateFunc func(interface{}) bool

func (pf PredicateFunc) Accept(v interface{}) bool {
	return pf(v)
}

// And returns a predicate that is true if all of the passed predicate are true for the input.
func And(predicates ...Predicate) Predicate {
	return PredicateFunc(func(v interface{}) bool {
		for _, p := range predicates {
			if !p.Accept(v) {
				return false
			}
		}
		return true
	})
}

// Or returns a predicate that is true if any of the passed predicate are true.  Furthermore, it
// stops executing predicate after the first true one.
func Or(predicates ...Predicate) Predicate {
	return PredicateFunc(func(v interface{}) bool {
		for _, p := range predicates {
			if p.Accept(v) {
				return true
			}
		}
		return false
	})
}

// Not returns a predicate that negates the condition defined by the passed predicate.
func Not(predicate Predicate) Predicate {
	return PredicateFunc(func(v interface{}) bool {
		return !predicate.Accept(v)
	})
}

// TruePredicate is a predicate that returns true for all inputs.
func True() Predicate {
	return PredicateFunc(func(v interface{}) bool { return true })
}

// FalsePredicate is a predicate that returns false for all inputs.
func False() Predicate {
	return PredicateFunc(func(v interface{}) bool { return false })
}

// A predicate that returns true if the value passed is a string and is equal to the value of 'value'
func StringEquals(value string) Predicate {
	return PredicateFunc(func(s interface{}) bool {
		return s.(string) == value
	})
}

// A predicate that returns true if the value passed contains a substring matching 'value'.
func StringContains(value string) Predicate {
	return PredicateFunc(func(s interface{}) bool {
		return strings.Contains(s.(string), value)
	})
}

// A predicate that returns true if the value passed starts with a substring matching 'value'.
func StringStartsWith(value string) Predicate {
	return PredicateFunc(func(s interface{}) bool {
		return strings.HasPrefix(s.(string), value)
	})
}

// A predicate that returns true if the value passed ends with a substring matching 'value'.
func StringEndsWith(value string) Predicate {
	return PredicateFunc(func(s interface{}) bool {
		return strings.HasSuffix(s.(string), value)
	})
}

// A predicate that returns true if the regex matches 'value'.
func StringMatches(regex *regexp.Regexp) Predicate {
	return PredicateFunc(func(s interface{}) bool {
		return regex.MatchString(s.(string))
	})
}

// ExtractedValueAccepted returns A predicate that extracts a value using the Extractor and passes that
// value to the provided predicate
func ExtractedValueAccepted(extractor Extractor, predicate Predicate) Predicate {
	return PredicateFunc(func(v interface{}) bool {
		return predicate.Accept(extractor.Extract(v))
	})
}

// PathMatches returns a predicate that returns true if the path matches the pathRegex.
func PathMatches(pathRegex *regexp.Regexp) Predicate {
	return ExtractedValueAccepted(ExtractPath(), StringMatches(pathRegex))
}

// PathEquals returns a predicate that returns true if the path equals 'path'
func PathEquals(path string) Predicate {
	return ExtractedValueAccepted(ExtractPath(), StringEquals(path))
}

// PathStartsWith returns a predicate that returns true if the path starts with 'path'
func PathStartsWith(path string) Predicate {
	return ExtractedValueAccepted(ExtractPath(), StringStartsWith(path))
}

// HeaderMatches returns a predicate that returns true if the header named 'name' matches 'regex'
func HeaderMatches(name string, regex *regexp.Regexp) Predicate {
	return ExtractedValueAccepted(ExtractHeader(name), StringMatches(regex))
}

// HeaderEquals returns a predicate that returns true if the header named 'name' equals 'value'
func HeaderEquals(name string, value string) Predicate {
	return ExtractedValueAccepted(ExtractHeader(name), StringEquals(value))
}

// HeaderEqualsIgnoreCase returns a predicate that returns true if if the header named 'name' equals 'value', ignoring case.
func HeaderEqualsIgnoreCase(name string, path string) Predicate {
	return ExtractedValueAccepted(UpperCaseExtractor(ExtractHeader(name)), StringEquals(strings.ToUpper(path)))
}

func HeaderContains(name string, path string) Predicate {
	return ExtractedValueAccepted(ExtractHeader(name), StringContains(path))
}

func HeaderContainsIgnoreCase(name string, path string) Predicate {
	return ExtractedValueAccepted(UpperCaseExtractor(ExtractHeader(name)), StringContains(strings.ToUpper(path)))
}

func HeaderStartsWith(name string, path string) Predicate {
	return ExtractedValueAccepted(ExtractHeader(name), StringStartsWith(path))
}

func RequestURIMatches(pathRegex *regexp.Regexp) Predicate {
	return ExtractedValueAccepted(ExtractRequestURI(), StringMatches(pathRegex))
}

func RequestURIEquals(path string) Predicate {
	return ExtractedValueAccepted(ExtractRequestURI(), StringEquals(path))
}

func RequestURIStartsWith(path string) Predicate {
	return ExtractedValueAccepted(ExtractRequestURI(), StringStartsWith(path))
}

// MethodIs returns a predicate that takes a request, extracts the method, and returns true if it equals the method
// provided, ignoring case.
func MethodIs(method string) Predicate {
	return ExtractedValueAccepted(UpperCaseExtractor(ExtractMethod()), StringEquals(strings.ToUpper(method)))
}

// QueryParamContainsIgnoreCase returns a Predicate that takes a request, extracts the query parameter specified and
// returns true if it equals the value provided.
func QueryParamEquals(name, value string) Predicate {
	return ExtractedValueAccepted(ExtractQueryParameter(name), StringEquals(value))
}

// QueryParamContainsIgnoreCase returns a Predicate that takes a request, extracts the query parameter specified and
// returns true if it equals the value provided, ignoring case.
func QueryParamEqualsIgnoreCase(name, value string) Predicate {
	return ExtractedValueAccepted(UpperCaseExtractor(ExtractQueryParameter(name)), StringEquals(strings.ToUpper(value)))
}

// QueryParamContainsIgnoreCase returns a Predicate that takes a request, extracts the query parameter specified and
// returns true if it contains the value provided.
func QueryParamContains(name, value string) Predicate {
	return ExtractedValueAccepted(ExtractQueryParameter(name), StringContains(value))
}

// QueryParamContainsIgnoreCase returns a Predicate that takes a request, extracts the query parameter specified and
// returns true if it contains the value provided, ignoring case.
func QueryParamContainsIgnoreCase(name, value string) Predicate {
	return ExtractedValueAccepted(UpperCaseExtractor(ExtractQueryParameter(name)), StringContains(strings.ToUpper(value)))
}

// QueryParamMatches returns a Predicate that takes a request, extracts the query parameter specified and
// returns true if the value matches the pattern provided.
func QueryParamMatches(name string, pattern *regexp.Regexp) Predicate {
	return ExtractedValueAccepted(ExtractQueryParameter(name), StringMatches(pattern))
}

// QueryParamStartsWith returns a Predicate that takes a request, extracts the query parameter specified and
// returns true if the value starts with the prefix provided.
func QueryParamStartsWith(name string, prefix string) Predicate {
	return ExtractedValueAccepted(ExtractQueryParameter(name), StringStartsWith(prefix))
}
