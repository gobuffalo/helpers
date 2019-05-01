package adler32h

import (
	"hash/adler32"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ChecksumKey = "Checksum"
)

func New() hctx.Map {
	return hctx.Map{

		ChecksumKey: Checksum,
	}
}

// Checksum returns the Adler-32 checksum of data.
var Checksum = adler32.Checksum
