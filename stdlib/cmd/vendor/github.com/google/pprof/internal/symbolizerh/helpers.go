package symbolizerh

import (
	"cmd/vendor/github.com/google/pprof/internal/symbolizer"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	DemangleKey = "Demangle"
)

func New() hctx.Map {
	return hctx.Map{

		DemangleKey: Demangle,
	}
}

// Demangle updates the function names in a profile with demangled C++
// names, simplified according to demanglerMode. If force is set,
// overwrite any names that appear already demangled.
var Demangle = symbolizer.Demangle
