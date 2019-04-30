package subtleh

import (
	"crypto/internal/subtle"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AnyOverlapKey = "AnyOverlap"

	AnyOverlapKey = "AnyOverlap"

	InexactOverlapKey = "InexactOverlap"

	InexactOverlapKey = "InexactOverlap"
)

func New() hctx.Map {
	return hctx.Map{

		AnyOverlapKey: AnyOverlap,

		AnyOverlapKey: AnyOverlap,

		InexactOverlapKey: InexactOverlap,

		InexactOverlapKey: InexactOverlap,
	}
}

// AnyOverlap reports whether x and y share memory at any (not necessarily
// corresponding) index. The memory beyond the slice length is ignored.
var AnyOverlap = subtle.AnyOverlap

// AnyOverlap reports whether x and y share memory at any (not necessarily
// corresponding) index. The memory beyond the slice length is ignored.
var AnyOverlap = subtle.AnyOverlap

// InexactOverlap reports whether x and y share memory at any non-corresponding
// index. The memory beyond the slice length is ignored. Note that x and y can
// have different lengths and still not have any inexact overlap.
//
// InexactOverlap can be used to implement the requirements of the crypto/cipher
// AEAD, Block, BlockMode and Stream interfaces.
var InexactOverlap = subtle.InexactOverlap

// InexactOverlap reports whether x and y share memory at any non-corresponding
// index. The memory beyond the slice length is ignored. Note that x and y can
// have different lengths and still not have any inexact overlap.
//
// InexactOverlap can be used to implement the requirements of the crypto/cipher
// AEAD, Block, BlockMode and Stream interfaces.
var InexactOverlap = subtle.InexactOverlap
