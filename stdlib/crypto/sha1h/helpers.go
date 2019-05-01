package sha1h

import (
	"crypto/sha1"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	SumKey = "Sum"
)

func New() hctx.Map {
	return hctx.Map{

		SumKey: Sum,
	}
}

// Sum returns the SHA-1 checksum of the data.
var Sum = sha1.Sum
