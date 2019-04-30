package ziph

import (
	"archive/zip"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	FileInfoHeaderKey = "FileInfoHeader"

	NewReaderKey = "NewReader"

	NewWriterKey = "NewWriter"

	OpenReaderKey = "OpenReader"

	RegisterCompressorKey = "RegisterCompressor"

	RegisterDecompressorKey = "RegisterDecompressor"
)

func New() hctx.Map {
	return hctx.Map{

		FileInfoHeaderKey: FileInfoHeader,

		NewReaderKey: NewReader,

		NewWriterKey: NewWriter,

		OpenReaderKey: OpenReader,

		RegisterCompressorKey: RegisterCompressor,

		RegisterDecompressorKey: RegisterDecompressor,
	}
}

// FileInfoHeader creates a partially-populated FileHeader from an
// os.FileInfo.
// Because os.FileInfo&#39;s Name method returns only the base name of
// the file it describes, it may be necessary to modify the Name field
// of the returned header to provide the full path name of the file.
// If compression is desired, callers should set the FileHeader.Method
// field; it is unset by default.
var FileInfoHeader = zip.FileInfoHeader

// NewReader returns a new Reader reading from r, which is assumed to
// have the given size in bytes.
var NewReader = zip.NewReader

// NewWriter returns a new Writer writing a zip file to w.
var NewWriter = zip.NewWriter

// OpenReader will open the Zip file specified by name and return a ReadCloser.
var OpenReader = zip.OpenReader

// RegisterCompressor registers custom compressors for a specified method ID.
// The common methods Store and Deflate are built in.
var RegisterCompressor = zip.RegisterCompressor

// RegisterDecompressor allows custom decompressors for a specified method ID.
// The common methods Store and Deflate are built in.
var RegisterDecompressor = zip.RegisterDecompressor
