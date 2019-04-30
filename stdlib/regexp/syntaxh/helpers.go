package syntaxh

import (
	"regexp/syntax"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	CompileKey = "Compile"

	EmptyOpContextKey = "EmptyOpContext"

	IsWordCharKey = "IsWordChar"

	ParseKey = "Parse"
)

func New() hctx.Map {
	return hctx.Map{

		CompileKey: Compile,

		EmptyOpContextKey: EmptyOpContext,

		IsWordCharKey: IsWordChar,

		ParseKey: Parse,
	}
}

// Compile compiles the regexp into a program to be executed.
// The regexp should have been simplified already (returned from re.Simplify).
var Compile = syntax.Compile

// EmptyOpContext returns the zero-width assertions
// satisfied at the position between the runes r1 and r2.
// Passing r1 == -1 indicates that the position is
// at the beginning of the text.
// Passing r2 == -1 indicates that the position is
// at the end of the text.
var EmptyOpContext = syntax.EmptyOpContext

// IsWordChar reports whether r is consider a ``word character&#39;&#39;
// during the evaluation of the \b and \B zero-width assertions.
// These assertions are ASCII-only: the word characters are [A-Za-z0-9_].
var IsWordChar = syntax.IsWordChar

// Parse parses a regular expression string s, controlled by the specified
// Flags, and returns a regular expression parse tree. The syntax is
// described in the top-level comment.
var Parse = syntax.Parse
