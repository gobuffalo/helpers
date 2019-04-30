package analysisutilh

import (
	"cmd/vendor/golang.org/x/tools/go/analysis/passes/internal/analysisutil"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	FormatKey = "Format"

	HasSideEffectsKey = "HasSideEffects"

	LineStartKey = "LineStart"

	ReadFileKey = "ReadFile"

	UnparenKey = "Unparen"
)

func New() hctx.Map {
	return hctx.Map{

		FormatKey: Format,

		HasSideEffectsKey: HasSideEffects,

		LineStartKey: LineStart,

		ReadFileKey: ReadFile,

		UnparenKey: Unparen,
	}
}

// Format returns a string representation of the expression.
var Format = analysisutil.Format

// HasSideEffects reports whether evaluation of e has side effects.
var HasSideEffects = analysisutil.HasSideEffects

// LineStart returns the position of the start of the specified line
// within file f, or NoPos if there is no line of that number.
var LineStart = analysisutil.LineStart

// ReadFile reads a file and adds it to the FileSet
// so that we can report errors against it using lineStart.
var ReadFile = analysisutil.ReadFile

// Unparen returns e with any enclosing parentheses stripped.
var Unparen = analysisutil.Unparen
