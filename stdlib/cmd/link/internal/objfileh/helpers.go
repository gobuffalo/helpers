package objfileh

import (
	"cmd/link/internal/objfile"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	LoadKey = "Load"
)

func New() hctx.Map {
	return hctx.Map{

		LoadKey: Load,
	}
}

// Load loads an object file f into library lib.
// The symbols loaded are added to syms.
var Load = objfile.Load
