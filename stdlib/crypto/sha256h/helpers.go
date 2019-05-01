package sha256h

import (
	"crypto/sha256"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	New224Key = "New224"

	Sum224Key = "Sum224"

	Sum256Key = "Sum256"
)

func New() hctx.Map {
	return hctx.Map{

		New224Key: New224,

		Sum224Key: Sum224,

		Sum256Key: Sum256,
	}
}

// New224 returns a new hash.Hash computing the SHA224 checksum.
var New224 = sha256.New224

// Sum224 returns the SHA224 checksum of the data.
var Sum224 = sha256.Sum224

// Sum256 returns the SHA256 checksum of the data.
var Sum256 = sha256.Sum256
