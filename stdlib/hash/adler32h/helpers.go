package adler32h

import (
	"hash/adler32"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ChecksumKey = "Checksum"

	NewKey = "New"
)

func New() hctx.Map {
	return hctx.Map{

		ChecksumKey: Checksum,

		NewKey: New,
	}
}

// Checksum returns the Adler-32 checksum of data.
var Checksum = adler32.Checksum

// New returns a new hash.Hash32 computing the Adler-32 checksum. Its
// Sum method will lay the value out in big-endian byte order. The
// returned Hash32 also implements encoding.BinaryMarshaler and
// encoding.BinaryUnmarshaler to marshal and unmarshal the internal
// state of the hash.
var New = adler32.New
