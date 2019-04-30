package syntaxh

import (
	"cmd/compile/internal/syntax"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	FdumpKey = "Fdump"

	FprintKey = "Fprint"

	MakePosKey = "MakePos"

	NewFileBaseKey = "NewFileBase"

	NewLineBaseKey = "NewLineBase"

	ParseKey = "Parse"

	ParseFileKey = "ParseFile"

	StringKey = "String"
)

func New() hctx.Map {
	return hctx.Map{

		FdumpKey: Fdump,

		FprintKey: Fprint,

		MakePosKey: MakePos,

		NewFileBaseKey: NewFileBase,

		NewLineBaseKey: NewLineBase,

		ParseKey: Parse,

		ParseFileKey: ParseFile,

		StringKey: String,
	}
}

// Fdump dumps the structure of the syntax tree rooted at n to w.
// It is intended for debugging purposes; no specific output format
// is guaranteed.
var Fdump = syntax.Fdump

var Fprint = syntax.Fprint

// MakePos returns a new Pos for the given PosBase, line and column.
var MakePos = syntax.MakePos

// NewFileBase returns a new PosBase for the given filename.
// A file PosBase&#39;s position is relative to itself, with the
// position being filename:1:1.
var NewFileBase = syntax.NewFileBase

// NewLineBase returns a new PosBase for a line directive &#34;line filename:line:col&#34;
// relative to pos, which is the position of the character immediately following
// the comment containing the line directive. For a directive in a line comment,
// that position is the beginning of the next line (i.e., the newline character
// belongs to the line comment).
var NewLineBase = syntax.NewLineBase

// Parse parses a single Go source file from src and returns the corresponding
// syntax tree. If there are errors, Parse will return the first error found,
// and a possibly partially constructed syntax tree, or nil.
//
// If errh != nil, it is called with each error encountered, and Parse will
// process as much source as possible. In this case, the returned syntax tree
// is only nil if no correct package clause was found.
// If errh is nil, Parse will terminate immediately upon encountering the first
// error, and the returned syntax tree is nil.
//
// If pragh != nil, it is called with each pragma encountered.
var Parse = syntax.Parse

// ParseFile behaves like Parse but it reads the source from the named file.
var ParseFile = syntax.ParseFile

var String = syntax.String
