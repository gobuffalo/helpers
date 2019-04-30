package cryptoh

import (
	"crypto"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	RegisterHashKey = "RegisterHash"
)

func New() hctx.Map {
	return hctx.Map{

		RegisterHashKey: RegisterHash,
	}
}

// RegisterHash registers a function that returns a new instance of the given
// hash function. This is intended to be called from the init function in
// packages that implement hash functions.
var RegisterHash = crypto.RegisterHash
