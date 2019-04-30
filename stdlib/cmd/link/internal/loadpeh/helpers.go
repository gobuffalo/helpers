package loadpeh

import (
	"cmd/link/internal/loadpe"

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

// Load loads the PE file pn from input.
// Symbols are written into syms, and a slice of the text symbols is returned.
// If an .rsrc section is found, its symbol is returned as rsrc.
var Load = loadpe.Load
