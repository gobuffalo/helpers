package lockedfileh

import (
	"cmd/go/internal/lockedfile"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	CreateKey = "Create"

	EditKey = "Edit"

	MutexAtKey = "MutexAt"

	OpenKey = "Open"

	OpenFileKey = "OpenFile"

	ReadKey = "Read"

	WriteKey = "Write"
)

func New() hctx.Map {
	return hctx.Map{

		CreateKey: Create,

		EditKey: Edit,

		MutexAtKey: MutexAt,

		OpenKey: Open,

		OpenFileKey: OpenFile,

		ReadKey: Read,

		WriteKey: Write,
	}
}

// Create is like os.Create, but returns a write-locked file.
var Create = lockedfile.Create

// Edit creates the named file with mode 0666 (before umask),
// but does not truncate existing contents.
//
// If Edit succeeds, methods on the returned File can be used for I/O.
// The associated file descriptor has mode O_RDWR and the file is write-locked.
var Edit = lockedfile.Edit

// MutexAt returns a new Mutex with Path set to the given non-empty path.
var MutexAt = lockedfile.MutexAt

// Open is like os.Open, but returns a read-locked file.
var Open = lockedfile.Open

// OpenFile is like os.OpenFile, but returns a locked file.
// If flag includes os.O_WRONLY or os.O_RDWR, the file is write-locked;
// otherwise, it is read-locked.
var OpenFile = lockedfile.OpenFile

// Read opens the named file with a read-lock and returns its contents.
var Read = lockedfile.Read

// Write opens the named file (creating it with the given permissions if needed),
// then write-locks it and overwrites it with the given content.
var Write = lockedfile.Write
