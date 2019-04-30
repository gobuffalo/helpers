package hmach

import (
	"crypto/hmac"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	EqualKey = "Equal"

	NewKey = "New"
)

func New() hctx.Map {
	return hctx.Map{

		EqualKey: Equal,

		NewKey: New,
	}
}

// Equal compares two MACs for equality without leaking timing information.
var Equal = hmac.Equal

// New returns a new HMAC hash using the given hash.Hash type and key.
// Note that unlike other hash implementations in the standard library,
// the returned Hash does not implement encoding.BinaryMarshaler
// or encoding.BinaryUnmarshaler.
var New = hmac.New
