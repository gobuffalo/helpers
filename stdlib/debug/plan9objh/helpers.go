package plan9objh

import (
	"debug/plan9obj"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewFileKey = "NewFile"

	OpenKey = "Open"
)

func New() hctx.Map {
	return hctx.Map{

		NewFileKey: NewFile,

		OpenKey: Open,
	}
}

// NewFile creates a new File for accessing a Plan 9 binary in an underlying reader.
// The Plan 9 binary is expected to start at position 0 in the ReaderAt.
var NewFile = plan9obj.NewFile

// Open opens the named file using os.Open and prepares it for use as a Plan 9 a.out binary.
var Open = plan9obj.Open
