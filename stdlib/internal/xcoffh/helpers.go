package xcoffh

import (
	"internal/xcoff"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewArchiveKey = "NewArchive"

	NewFileKey = "NewFile"

	OpenKey = "Open"

	OpenArchiveKey = "OpenArchive"
)

func New() hctx.Map {
	return hctx.Map{

		NewArchiveKey: NewArchive,

		NewFileKey: NewFile,

		OpenKey: Open,

		OpenArchiveKey: OpenArchive,
	}
}

// NewArchive creates a new Archive for accessing an AIX big archive in an underlying reader.
var NewArchive = xcoff.NewArchive

// NewFile creates a new File for accessing an XCOFF binary in an underlying reader.
var NewFile = xcoff.NewFile

// Open opens the named file using os.Open and prepares it for use as an XCOFF binary.
var Open = xcoff.Open

// OpenArchive opens the named archive using os.Open and prepares it for use
// as an AIX big archive.
var OpenArchive = xcoff.OpenArchive
