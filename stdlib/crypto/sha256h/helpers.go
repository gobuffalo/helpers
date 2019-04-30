package sha256h

import (
	"crypto/sha256"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewKey = "New"

	New224Key = "New224"

	Sum224Key = "Sum224"

	Sum256Key = "Sum256"
)

func New() hctx.Map {
	return hctx.Map{

		NewKey: New,

		New224Key: New224,

		Sum224Key: Sum224,

		Sum256Key: Sum256,
	}
}

// New returns a new hash.Hash computing the SHA256 checksum. The Hash
// also implements encoding.BinaryMarshaler and
// encoding.BinaryUnmarshaler to marshal and unmarshal the internal
// state of the hash.
var New = sha256.New

// New224 returns a new hash.Hash computing the SHA224 checksum.
var New224 = sha256.New224

// Sum224 returns the SHA224 checksum of the data.
var Sum224 = sha256.Sum224

// Sum256 returns the SHA256 checksum of the data.
var Sum256 = sha256.Sum256
