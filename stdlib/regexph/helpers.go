package regexph

import (
	"regexp"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	CompileKey = "Compile"

	CompilePOSIXKey = "CompilePOSIX"

	MatchKey = "Match"

	MatchReaderKey = "MatchReader"

	MatchStringKey = "MatchString"

	MustCompileKey = "MustCompile"

	MustCompilePOSIXKey = "MustCompilePOSIX"

	QuoteMetaKey = "QuoteMeta"
)

func New() hctx.Map {
	return hctx.Map{

		CompileKey: Compile,

		CompilePOSIXKey: CompilePOSIX,

		MatchKey: Match,

		MatchReaderKey: MatchReader,

		MatchStringKey: MatchString,

		MustCompileKey: MustCompile,

		MustCompilePOSIXKey: MustCompilePOSIX,

		QuoteMetaKey: QuoteMeta,
	}
}

// Compile parses a regular expression and returns, if successful,
// a Regexp object that can be used to match against text.
//
// When matching against text, the regexp returns a match that
// begins as early as possible in the input (leftmost), and among those
// it chooses the one that a backtracking search would have found first.
// This so-called leftmost-first matching is the same semantics
// that Perl, Python, and other implementations use, although this
// package implements it without the expense of backtracking.
// For POSIX leftmost-longest matching, see CompilePOSIX.
var Compile = regexp.Compile

// CompilePOSIX is like Compile but restricts the regular expression
// to POSIX ERE (egrep) syntax and changes the match semantics to
// leftmost-longest.
//
// That is, when matching against text, the regexp returns a match that
// begins as early as possible in the input (leftmost), and among those
// it chooses a match that is as long as possible.
// This so-called leftmost-longest matching is the same semantics
// that early regular expression implementations used and that POSIX
// specifies.
//
// However, there can be multiple leftmost-longest matches, with different
// submatch choices, and here this package diverges from POSIX.
// Among the possible leftmost-longest matches, this package chooses
// the one that a backtracking search would have found first, while POSIX
// specifies that the match be chosen to maximize the length of the first
// subexpression, then the second, and so on from left to right.
// The POSIX rule is computationally prohibitive and not even well-defined.
// See http://swtch.com/~rsc/regexp/regexp2.html#posix for details.
var CompilePOSIX = regexp.CompilePOSIX

// Match checks whether a textual regular expression
// matches a byte slice. More complicated queries need
// to use Compile and the full Regexp interface.
var Match = regexp.Match

// MatchReader checks whether a textual regular expression matches the text
// read by the RuneReader. More complicated queries need to use Compile and
// the full Regexp interface.
var MatchReader = regexp.MatchReader

// MatchString checks whether a textual regular expression
// matches a string. More complicated queries need
// to use Compile and the full Regexp interface.
var MatchString = regexp.MatchString

// MustCompile is like Compile but panics if the expression cannot be parsed.
// It simplifies safe initialization of global variables holding compiled regular
// expressions.
var MustCompile = regexp.MustCompile

// MustCompilePOSIX is like CompilePOSIX but panics if the expression cannot be parsed.
// It simplifies safe initialization of global variables holding compiled regular
// expressions.
var MustCompilePOSIX = regexp.MustCompilePOSIX

// QuoteMeta returns a string that quotes all regular expression metacharacters
// inside the argument text; the returned string is a regular expression matching
// the literal text. For example, QuoteMeta(`[foo]`) returns `\[foo\]`.
var QuoteMeta = regexp.QuoteMeta
