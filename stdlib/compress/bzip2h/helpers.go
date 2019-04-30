package bzip2h

import (
	"compress/bzip2"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewReaderKey = "NewReader"
)

func New() hctx.Map {
	return hctx.Map{

		NewReaderKey: NewReader,
	}
}

// NewReader returns an io.Reader which decompresses bzip2 data from r.
// If r does not also implement io.ByteReader,
// the decompressor may read more data than necessary from r.
var NewReader = bzip2.NewReader
