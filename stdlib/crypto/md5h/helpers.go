package md5h

import (
	"crypto/md5"

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

// Sum returns the MD5 checksum of the data.
var Sum = md5.Sum
