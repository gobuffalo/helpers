package loadmachoh

import (
	"cmd/link/internal/loadmacho"

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

// Load loads the Mach-O file pn from f.
// Symbols are written into syms, and a slice of the text symbols is returned.
var Load = loadmacho.Load
