package sha1h

import (
	"crypto/sha1"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewKey = "New"

	SumKey = "Sum"
)

func New() hctx.Map {
	return hctx.Map{

		NewKey: New,

		SumKey: Sum,
	}
}

// New returns a new hash.Hash computing the SHA1 checksum. The Hash also
// implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to
// marshal and unmarshal the internal state of the hash.
var New = sha1.New

// Sum returns the SHA-1 checksum of the data.
var Sum = sha1.Sum
