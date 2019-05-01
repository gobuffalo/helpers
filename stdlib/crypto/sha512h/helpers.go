package sha512h

import (
	"crypto/sha512"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	New384Key = "New384"

	New512_224Key = "New512_224"

	New512_256Key = "New512_256"

	Sum384Key = "Sum384"

	Sum512Key = "Sum512"

	Sum512_224Key = "Sum512_224"

	Sum512_256Key = "Sum512_256"
)

func New() hctx.Map {
	return hctx.Map{

		New384Key: New384,

		New512_224Key: New512_224,

		New512_256Key: New512_256,

		Sum384Key: Sum384,

		Sum512Key: Sum512,

		Sum512_224Key: Sum512_224,

		Sum512_256Key: Sum512_256,
	}
}

// New384 returns a new hash.Hash computing the SHA-384 checksum.
var New384 = sha512.New384

// New512_224 returns a new hash.Hash computing the SHA-512/224 checksum.
var New512_224 = sha512.New512_224

// New512_256 returns a new hash.Hash computing the SHA-512/256 checksum.
var New512_256 = sha512.New512_256

// Sum384 returns the SHA384 checksum of the data.
var Sum384 = sha512.Sum384

// Sum512 returns the SHA512 checksum of the data.
var Sum512 = sha512.Sum512

// Sum512_224 returns the Sum512/224 checksum of the data.
var Sum512_224 = sha512.Sum512_224

// Sum512_256 returns the Sum512/256 checksum of the data.
var Sum512_256 = sha512.Sum512_256
