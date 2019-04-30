package lzwh

import (
	"compress/lzw"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewReaderKey = "NewReader"

	NewWriterKey = "NewWriter"
)

func New() hctx.Map {
	return hctx.Map{

		NewReaderKey: NewReader,

		NewWriterKey: NewWriter,
	}
}

// NewReader creates a new io.ReadCloser.
// Reads from the returned io.ReadCloser read and decompress data from r.
// If r does not also implement io.ByteReader,
// the decompressor may read more data than necessary from r.
// It is the caller&#39;s responsibility to call Close on the ReadCloser when
// finished reading.
// The number of bits to use for literal codes, litWidth, must be in the
// range [2,8] and is typically 8. It must equal the litWidth
// used during compression.
var NewReader = lzw.NewReader

// NewWriter creates a new io.WriteCloser.
// Writes to the returned io.WriteCloser are compressed and written to w.
// It is the caller&#39;s responsibility to call Close on the WriteCloser when
// finished writing.
// The number of bits to use for literal codes, litWidth, must be in the
// range [2,8] and is typically 8. Input bytes must be less than 1&lt;&lt;litWidth.
var NewWriter = lzw.NewWriter
