package predicate

import (
	. "go-http-matchers/extractor"
	"strings"
)

// MethodIs returns a predicate that takes a request, extracts the method, and returns true if it equals the method
// provided, ignoring case.
func MethodIs(method string) Predicate {
	return ExtractedValueAccepted(UpperCaseExtractor(ExtractMethod()), StringEquals(strings.ToUpper(method)))
}
