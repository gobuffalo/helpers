package cpuh

import (
	"internal/cpu"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	InitializeKey = "Initialize"
)

func New() hctx.Map {
	return hctx.Map{

		InitializeKey: Initialize,
	}
}

// Initialize examines the processor and sets the relevant variables above.
// This is called by the runtime package early in program initialization,
// before normal init functions are run. env is set by runtime if the OS supports
// cpu feature options in GODEBUG.
var Initialize = cpu.Initialize
