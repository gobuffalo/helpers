package tarh

import (
	"archive/tar"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	FileInfoHeaderKey = "FileInfoHeader"

	NewReaderKey = "NewReader"

	NewWriterKey = "NewWriter"
)

func New() hctx.Map {
	return hctx.Map{

		FileInfoHeaderKey: FileInfoHeader,

		NewReaderKey: NewReader,

		NewWriterKey: NewWriter,
	}
}

// FileInfoHeader creates a partially-populated Header from fi.
// If fi describes a symlink, FileInfoHeader records link as the link target.
// If fi describes a directory, a slash is appended to the name.
//
// Since os.FileInfo&#39;s Name method only returns the base name of
// the file it describes, it may be necessary to modify Header.Name
// to provide the full path name of the file.
var FileInfoHeader = tar.FileInfoHeader

// NewReader creates a new Reader reading from r.
var NewReader = tar.NewReader

// NewWriter creates a new Writer writing to w.
var NewWriter = tar.NewWriter
