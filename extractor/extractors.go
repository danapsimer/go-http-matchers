package extractor

import (
	"gopkg.in/xmlpath.v2"
	"net/http"
	"strings"
)

// Extractor can extract a value from another value by calling the Extract method.
type Extractor interface {
	Extract(interface{}) interface{}
}

// ExtractorFunc is a function that calls itself when it's Extract method is called.
type ExtractorFunc func(interface{}) interface{}

func (ef ExtractorFunc) Extract(v interface{}) interface{} {
	return ef(v)
}

// RequestKeyIdentity is a Extractor that returns the entire Request.
func IdentityExtractor() Extractor {
	return ExtractorFunc(func(r interface{}) interface{} {
		return r
	})
}

// MethodExtractor is an extractor that returns the method of an http.Request.
func ExtractMethod() Extractor {
	return ExtractorFunc(func(r interface{}) interface{} {
		return r.(*http.Request).Method
	})
}

// ExtractPath is an Extractor that returns the request URL's Path property.  This is just path path portion of the URI.
func ExtractPath() Extractor {
	return ExtractorFunc(func(r interface{}) interface{} {
		return r.(*http.Request).URL.Path
	})
}

// ExtractRequestURI is an Extractor that returns the request URL's RequestURI property.  This is the path and the query
// portions of the URI.
func ExtractRequestURI() Extractor {
	return ExtractorFunc(func(r interface{}) interface{} {
		return r.(*http.Request).URL.RequestURI()
	})
}

// ExtractHeader returns a Extractor that returns the value of the header named 'name'
func ExtractHeader(name string) Extractor {
	return ExtractorFunc(func(r interface{}) interface{} {
		return r.(*http.Request).Header.Get(name)
	})
}

// UpperCaseExtractor returns an Extractor that decorates the passed extractor by applying strings.ToUpper to the
// value returned.
func UpperCaseExtractor(extractor Extractor) Extractor {
	return ExtractorFunc(func(v interface{}) interface{} {
		value := extractor.Extract(v)
		if value == nil {
			return nil
		}
		return strings.ToUpper(value.(string))
	})
}

// ExtractXPathString returns a Extractor that uses XPATH expression to extract a string from the Body of the
// Request.
func ExtractXPathString(xpath string) Extractor {
	path := xmlpath.MustCompile(xpath)
	return ExtractorFunc(func(r interface{}) interface{} {
		str := ""
		root, err := xmlpath.Parse(r.(*http.Request).Body)
		if err == nil {
			str, _ = path.String(root)
		}
		return str
	})
}

// ExtractPathElementByIndex returns a Extractor that extracts the path element at the given position.  A
// negative number denotes a position from the end (starting at 1 e.g. -1 is the last element in the path).  For
// positive inputs, the counting starts at 1 as well.
func ExtractPathElementByIndex(idx int) Extractor {
	return ExtractorFunc(func(r interface{}) interface{} {
		elements := strings.Split(r.(*http.Request).URL.Path, "/")
		var i int
		if idx < 0 {
			i = len(elements) + idx
		} else {
			i = idx
		}
		if i < 0 || i >= len(elements) {
			return ""
		}
		return elements[i]
	})
}

// ExtractQueryParameter returns a Extractor that extracts a query parameters value.
func ExtractQueryParameter(name string) Extractor {
	return ExtractorFunc(func(r interface{}) interface{} {
		return r.(*http.Request).URL.Query().Get(name)
	})
}
