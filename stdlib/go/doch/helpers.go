package doch

import (
	"go/doc"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ExamplesKey = "Examples"

	IsPredeclaredKey = "IsPredeclared"

	SynopsisKey = "Synopsis"

	ToHTMLKey = "ToHTML"

	ToTextKey = "ToText"
)

func New() hctx.Map {
	return hctx.Map{

		ExamplesKey: Examples,

		IsPredeclaredKey: IsPredeclared,

		SynopsisKey: Synopsis,

		ToHTMLKey: ToHTML,

		ToTextKey: ToText,
	}
}

// Examples returns the examples found in the files, sorted by Name field.
// The Order fields record the order in which the examples were encountered.
//
// Playable Examples must be in a package whose name ends in &#34;_test&#34;.
// An Example is &#34;playable&#34; (the Play field is non-nil) in either of these
// circumstances:
//   - The example function is self-contained: the function references only
//     identifiers from other packages (or predeclared identifiers, such as
//     &#34;int&#34;) and the test file does not include a dot import.
//   - The entire test file is the example: the file contains exactly one
//     example function, zero test or benchmark functions, and at least one
//     top-level function, type, variable, or constant declaration other
//     than the example function.
var Examples = doc.Examples

// IsPredeclared reports whether s is a predeclared identifier.
var IsPredeclared = doc.IsPredeclared

// Synopsis returns a cleaned version of the first sentence in s.
// That sentence ends after the first period followed by space and
// not preceded by exactly one uppercase letter. The result string
// has no \n, \r, or \t characters and uses only single spaces between
// words. If s starts with any of the IllegalPrefixes, the result
// is the empty string.
var Synopsis = doc.Synopsis

// ToHTML converts comment text to formatted HTML.
// The comment was prepared by DocReader,
// so it is known not to have leading, trailing blank lines
// nor to have trailing spaces at the end of lines.
// The comment markers have already been removed.
//
// Each span of unindented non-blank lines is converted into
// a single paragraph. There is one exception to the rule: a span that
// consists of a single line, is followed by another paragraph span,
// begins with a capital letter, and contains no punctuation
// other than parentheses and commas is formatted as a heading.
//
// A span of indented lines is converted into a &lt;pre&gt; block,
// with the common indent prefix removed.
//
// URLs in the comment text are converted into links; if the URL also appears
// in the words map, the link is taken from the map (if the corresponding map
// value is the empty string, the URL is not converted into a link).
//
// Go identifiers that appear in the words map are italicized; if the corresponding
// map value is not the empty string, it is considered a URL and the word is converted
// into a link.
var ToHTML = doc.ToHTML

// ToText prepares comment text for presentation in textual output.
// It wraps paragraphs of text to width or fewer Unicode code points
// and then prefixes each line with the indent. In preformatted sections
// (such as program text), it prefixes each non-blank line with preIndent.
var ToText = doc.ToText
