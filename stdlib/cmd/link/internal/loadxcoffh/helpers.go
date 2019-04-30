package loadxcoffh

import (
	"cmd/link/internal/loadxcoff"

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

// Load loads the Xcoff file pn from f.
// Symbols are written into syms, and a slice of the text symbols is returned.
var Load = loadxcoff.Load
