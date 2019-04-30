package lexh

import (
	"cmd/asm/internal/lex"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	IsRegisterShiftKey = "IsRegisterShift"

	MakeKey = "Make"

	NewInputKey = "NewInput"

	NewLexerKey = "NewLexer"

	NewSliceKey = "NewSlice"

	NewTokenizerKey = "NewTokenizer"

	TokenizeKey = "Tokenize"
)

func New() hctx.Map {
	return hctx.Map{

		IsRegisterShiftKey: IsRegisterShift,

		MakeKey: Make,

		NewInputKey: NewInput,

		NewLexerKey: NewLexer,

		NewSliceKey: NewSlice,

		NewTokenizerKey: NewTokenizer,

		TokenizeKey: Tokenize,
	}
}

// IsRegisterShift reports whether the token is one of the ARM register shift operators.
var IsRegisterShift = lex.IsRegisterShift

// Make returns a Token with the given rune (ScanToken) and text representation.
var Make = lex.Make

// NewInput returns an Input from the given path.
var NewInput = lex.NewInput

// NewLexer returns a lexer for the named file and the given link context.
var NewLexer = lex.NewLexer

var NewSlice = lex.NewSlice

var NewTokenizer = lex.NewTokenizer

// Tokenize turns a string into a list of Tokens; used to parse the -D flag and in tests.
var Tokenize = lex.Tokenize
