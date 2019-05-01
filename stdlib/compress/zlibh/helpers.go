package zlibh

import (
	"compress/zlib"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	BestCompressionKey = "BestCompression"

	BestSpeedKey = "BestSpeed"

	DefaultCompressionKey = "DefaultCompression"

	HuffmanOnlyKey = "HuffmanOnly"

	NewReaderKey = "NewReader"

	NewReaderDictKey = "NewReaderDict"

	NewWriterKey = "NewWriter"

	NewWriterLevelKey = "NewWriterLevel"

	NewWriterLevelDictKey = "NewWriterLevelDict"

	NoCompressionKey = "NoCompression"
)

func New() hctx.Map {
	return hctx.Map{

		BestCompressionKey: BestCompression,

		BestSpeedKey: BestSpeed,

		DefaultCompressionKey: DefaultCompression,

		HuffmanOnlyKey: HuffmanOnly,

		NewReaderKey: NewReader,

		NewReaderDictKey: NewReaderDict,

		NewWriterKey: NewWriter,

		NewWriterLevelKey: NewWriterLevel,

		NewWriterLevelDictKey: NewWriterLevelDict,

		NoCompressionKey: NoCompression,
	}
}

var BestCompression = zlib.BestCompression

var BestSpeed = zlib.BestSpeed

var DefaultCompression = zlib.DefaultCompression

var HuffmanOnly = zlib.HuffmanOnly

// NewReader creates a new ReadCloser.
// Reads from the returned ReadCloser read and decompress data from r.
// If r does not implement io.ByteReader, the decompressor may read more
// data than necessary from r.
// It is the caller&#39;s responsibility to call Close on the ReadCloser when done.
//
// The ReadCloser returned by NewReader also implements Resetter.
var NewReader = zlib.NewReader

// NewReaderDict is like NewReader but uses a preset dictionary.
// NewReaderDict ignores the dictionary if the compressed data does not refer to it.
// If the compressed data refers to a different dictionary, NewReaderDict returns ErrDictionary.
//
// The ReadCloser returned by NewReaderDict also implements Resetter.
var NewReaderDict = zlib.NewReaderDict

// NewWriter creates a new Writer.
// Writes to the returned Writer are compressed and written to w.
//
// It is the caller&#39;s responsibility to call Close on the Writer when done.
// Writes may be buffered and not flushed until Close.
var NewWriter = zlib.NewWriter

// NewWriterLevel is like NewWriter but specifies the compression level instead
// of assuming DefaultCompression.
//
// The compression level can be DefaultCompression, NoCompression, HuffmanOnly
// or any integer value between BestSpeed and BestCompression inclusive.
// The error returned will be nil if the level is valid.
var NewWriterLevel = zlib.NewWriterLevel

// NewWriterLevelDict is like NewWriterLevel but specifies a dictionary to
// compress with.
//
// The dictionary may be nil. If not, its contents should not be modified until
// the Writer is closed.
var NewWriterLevelDict = zlib.NewWriterLevelDict

var NoCompression = zlib.NoCompression
