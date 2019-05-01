package gziph

import (
	"compress/gzip"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	BestCompressionKey = "BestCompression"

	BestSpeedKey = "BestSpeed"

	DefaultCompressionKey = "DefaultCompression"

	HuffmanOnlyKey = "HuffmanOnly"

	NewReaderKey = "NewReader"

	NewWriterKey = "NewWriter"

	NewWriterLevelKey = "NewWriterLevel"

	NoCompressionKey = "NoCompression"
)

func New() hctx.Map {
	return hctx.Map{

		BestCompressionKey: BestCompression,

		BestSpeedKey: BestSpeed,

		DefaultCompressionKey: DefaultCompression,

		HuffmanOnlyKey: HuffmanOnly,

		NewReaderKey: NewReader,

		NewWriterKey: NewWriter,

		NewWriterLevelKey: NewWriterLevel,

		NoCompressionKey: NoCompression,
	}
}

var BestCompression = gzip.BestCompression

var BestSpeed = gzip.BestSpeed

var DefaultCompression = gzip.DefaultCompression

var HuffmanOnly = gzip.HuffmanOnly

// NewReader creates a new Reader reading the given reader.
// If r does not also implement io.ByteReader,
// the decompressor may read more data than necessary from r.
//
// It is the caller&#39;s responsibility to call Close on the Reader when done.
//
// The Reader.Header fields will be valid in the Reader returned.
var NewReader = gzip.NewReader

// NewWriter returns a new Writer.
// Writes to the returned writer are compressed and written to w.
//
// It is the caller&#39;s responsibility to call Close on the Writer when done.
// Writes may be buffered and not flushed until Close.
//
// Callers that wish to set the fields in Writer.Header must do so before
// the first call to Write, Flush, or Close.
var NewWriter = gzip.NewWriter

// NewWriterLevel is like NewWriter but specifies the compression level instead
// of assuming DefaultCompression.
//
// The compression level can be DefaultCompression, NoCompression, HuffmanOnly
// or any integer value between BestSpeed and BestCompression inclusive.
// The error returned will be nil if the level is valid.
var NewWriterLevel = gzip.NewWriterLevel

var NoCompression = gzip.NoCompression
