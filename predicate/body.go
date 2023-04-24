package predicate

import (
	"github.com/bluesoftdev/go-http-matchers/extractor"
	"regexp"
	"strings"
)

// BodyXPathEquals checks to see if the result of the xpath expression, matches the string supplied in the 'value'
// parameter.
func BodyXPathEquals(xpath, value string) Predicate {
	return ExtractedValueAccepted(extractor.ExtractXPathString(xpath), StringEquals(value))
}

// BodyXPathEqualsIgnoreCase similar to BodyXPathEquals but ignores case when comparing the strings.
func BodyXPathEqualsIgnoreCase(xpath, value string) Predicate {
	return ExtractedValueAccepted(extractor.UpperCaseExtractor(extractor.ExtractXPathString(xpath)),
		StringEquals(strings.ToUpper(value)))
}

// BodyXPathMatches checks to see if the result of the xpath expression, matches the regular expression given in the
// 'pattern' parameter.
func BodyXPathMatches(xpath string, pattern *regexp.Regexp) Predicate {
	return ExtractedValueAccepted(extractor.ExtractXPathString(xpath), StringMatches(pattern))
}
