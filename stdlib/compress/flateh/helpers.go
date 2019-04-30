package flateh

import (
	"compress/flate"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewReaderKey = "NewReader"

	NewReaderDictKey = "NewReaderDict"

	NewWriterKey = "NewWriter"

	NewWriterDictKey = "NewWriterDict"
)

func New() hctx.Map {
	return hctx.Map{

		NewReaderKey: NewReader,

		NewReaderDictKey: NewReaderDict,

		NewWriterKey: NewWriter,

		NewWriterDictKey: NewWriterDict,
	}
}

// NewReader returns a new ReadCloser that can be used
// to read the uncompressed version of r.
// If r does not also implement io.ByteReader,
// the decompressor may read more data than necessary from r.
// It is the caller&#39;s responsibility to call Close on the ReadCloser
// when finished reading.
//
// The ReadCloser returned by NewReader also implements Resetter.
var NewReader = flate.NewReader

// NewReaderDict is like NewReader but initializes the reader
// with a preset dictionary. The returned Reader behaves as if
// the uncompressed data stream started with the given dictionary,
// which has already been read. NewReaderDict is typically used
// to read data compressed by NewWriterDict.
//
// The ReadCloser returned by NewReader also implements Resetter.
var NewReaderDict = flate.NewReaderDict

// NewWriter returns a new Writer compressing data at the given level.
// Following zlib, levels range from 1 (BestSpeed) to 9 (BestCompression);
// higher levels typically run slower but compress more. Level 0
// (NoCompression) does not attempt any compression; it only adds the
// necessary DEFLATE framing.
// Level -1 (DefaultCompression) uses the default compression level.
// Level -2 (HuffmanOnly) will use Huffman compression only, giving
// a very fast compression for all types of input, but sacrificing considerable
// compression efficiency.
//
// If level is in the range [-2, 9] then the error returned will be nil.
// Otherwise the error returned will be non-nil.
var NewWriter = flate.NewWriter

// NewWriterDict is like NewWriter but initializes the new
// Writer with a preset dictionary. The returned Writer behaves
// as if the dictionary had been written to it without producing
// any compressed output. The compressed data written to w
// can only be decompressed by a Reader initialized with the
// same dictionary.
var NewWriterDict = flate.NewWriterDict
