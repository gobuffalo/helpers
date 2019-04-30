package randutilh

import (
	"crypto/internal/randutil"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	MaybeReadByteKey = "MaybeReadByte"
)

func New() hctx.Map {
	return hctx.Map{

		MaybeReadByteKey: MaybeReadByte,
	}
}

// MaybeReadByte reads a single byte from r with ~50% probability. This is used
// to ensure that callers do not depend on non-guaranteed behaviour, e.g.
// assuming that rsa.GenerateKey is deterministic w.r.t. a given random stream.
//
// This does not affect tests that pass a stream of fixed bytes as the random
// source (e.g. a zeroReader).
var MaybeReadByte = randutil.MaybeReadByte
