package templateh

import (
	"html/template"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	HTMLEscapeKey = "HTMLEscape"

	HTMLEscapeStringKey = "HTMLEscapeString"

	HTMLEscaperKey = "HTMLEscaper"

	IsTrueKey = "IsTrue"

	JSEscapeKey = "JSEscape"

	JSEscapeStringKey = "JSEscapeString"

	JSEscaperKey = "JSEscaper"

	MustKey = "Must"

	ParseFilesKey = "ParseFiles"

	ParseGlobKey = "ParseGlob"

	URLQueryEscaperKey = "URLQueryEscaper"
)

func New() hctx.Map {
	return hctx.Map{

		HTMLEscapeKey: HTMLEscape,

		HTMLEscapeStringKey: HTMLEscapeString,

		HTMLEscaperKey: HTMLEscaper,

		IsTrueKey: IsTrue,

		JSEscapeKey: JSEscape,

		JSEscapeStringKey: JSEscapeString,

		JSEscaperKey: JSEscaper,

		MustKey: Must,

		ParseFilesKey: ParseFiles,

		ParseGlobKey: ParseGlob,

		URLQueryEscaperKey: URLQueryEscaper,
	}
}

// HTMLEscape writes to w the escaped HTML equivalent of the plain text data b.
var HTMLEscape = template.HTMLEscape

// HTMLEscapeString returns the escaped HTML equivalent of the plain text data s.
var HTMLEscapeString = template.HTMLEscapeString

// HTMLEscaper returns the escaped HTML equivalent of the textual
// representation of its arguments.
var HTMLEscaper = template.HTMLEscaper

// IsTrue reports whether the value is &#39;true&#39;, in the sense of not the zero of its type,
// and whether the value has a meaningful truth value. This is the definition of
// truth used by if and other such actions.
var IsTrue = template.IsTrue

// JSEscape writes to w the escaped JavaScript equivalent of the plain text data b.
var JSEscape = template.JSEscape

// JSEscapeString returns the escaped JavaScript equivalent of the plain text data s.
var JSEscapeString = template.JSEscapeString

// JSEscaper returns the escaped JavaScript equivalent of the textual
// representation of its arguments.
var JSEscaper = template.JSEscaper

// Must is a helper that wraps a call to a function returning (*Template, error)
// and panics if the error is non-nil. It is intended for use in variable initializations
// such as
// 	var t = template.Must(template.New(&#34;name&#34;).Parse(&#34;html&#34;))
var Must = template.Must

// ParseFiles creates a new Template and parses the template definitions from
// the named files. The returned template&#39;s name will have the (base) name and
// (parsed) contents of the first file. There must be at least one file.
// If an error occurs, parsing stops and the returned *Template is nil.
//
// When parsing multiple files with the same name in different directories,
// the last one mentioned will be the one that results.
// For instance, ParseFiles(&#34;a/foo&#34;, &#34;b/foo&#34;) stores &#34;b/foo&#34; as the template
// named &#34;foo&#34;, while &#34;a/foo&#34; is unavailable.
var ParseFiles = template.ParseFiles

// ParseGlob creates a new Template and parses the template definitions from the
// files identified by the pattern, which must match at least one file. The
// returned template will have the (base) name and (parsed) contents of the
// first file matched by the pattern. ParseGlob is equivalent to calling
// ParseFiles with the list of files matched by the pattern.
//
// When parsing multiple files with the same name in different directories,
// the last one mentioned will be the one that results.
var ParseGlob = template.ParseGlob

// URLQueryEscaper returns the escaped value of the textual representation of
// its arguments in a form suitable for embedding in a URL query.
var URLQueryEscaper = template.URLQueryEscaper
