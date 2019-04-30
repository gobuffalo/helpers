package mathh

import (
	"runtime/internal/math"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	MulUintptrKey = "MulUintptr"
)

func New() hctx.Map {
	return hctx.Map{

		MulUintptrKey: MulUintptr,
	}
}

// MulUintptr returns a * b and whether the multiplication overflowed.
// On supported platforms this is an intrinsic lowered by the compiler.
var MulUintptr = math.MulUintptr
