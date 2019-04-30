package loadelfh

import (
	"cmd/link/internal/loadelf"

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

// Load loads the ELF file pn from f.
// Symbols are written into syms, and a slice of the text symbols is returned.
//
// On ARM systems, Load will attempt to determine what ELF header flags to
// emit by scanning the attributes in the ELF file being loaded. The
// parameter initEhdrFlags contains the current header flags for the output
// object, and the returned ehdrFlags contains what this Load function computes.
// TODO: find a better place for this logic.
var Load = loadelf.Load
