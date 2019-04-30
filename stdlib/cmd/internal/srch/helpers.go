package srch

import (
	"cmd/internal/src"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	MakePosKey = "MakePos"

	NewFileBaseKey = "NewFileBase"

	NewInliningBaseKey = "NewInliningBase"

	NewLinePragmaBaseKey = "NewLinePragmaBase"
)

func New() hctx.Map {
	return hctx.Map{

		MakePosKey: MakePos,

		NewFileBaseKey: NewFileBase,

		NewInliningBaseKey: NewInliningBase,

		NewLinePragmaBaseKey: NewLinePragmaBase,
	}
}

// MakePos creates a new Pos value with the given base, and (file-absolute)
// line and column.
var MakePos = src.MakePos

// NewFileBase returns a new *PosBase for a file with the given (relative and
// absolute) filenames.
var NewFileBase = src.NewFileBase

// NewInliningBase returns a copy of the old PosBase with the given inlining
// index. If old == nil, the resulting PosBase has no filename.
var NewInliningBase = src.NewInliningBase

// NewLinePragmaBase returns a new *PosBase for a line directive of the form
//      //line filename:line:col
//      /*line filename:line:col*/
// at position pos.
var NewLinePragmaBase = src.NewLinePragmaBase
