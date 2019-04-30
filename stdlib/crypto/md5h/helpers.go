package md5h

import (
	"crypto/md5"

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

// New returns a new hash.Hash computing the MD5 checksum. The Hash also
// implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to
// marshal and unmarshal the internal state of the hash.
var New = md5.New

// Sum returns the MD5 checksum of the data.
var Sum = md5.Sum
