package sysdllh

import (
	"internal/syscall/windows/sysdll"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AddKey = "Add"
)

func New() hctx.Map {
	return hctx.Map{

		AddKey: Add,
	}
}

// Add notes that dll is a system32 DLL which should only be loaded
// from the Windows SYSTEM32 directory. It returns its argument back,
// for ease of use in generated code.
var Add = sysdll.Add
