package machoh

import (
	"debug/macho"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewFatFileKey = "NewFatFile"

	NewFileKey = "NewFile"

	OpenKey = "Open"

	OpenFatKey = "OpenFat"
)

func New() hctx.Map {
	return hctx.Map{

		NewFatFileKey: NewFatFile,

		NewFileKey: NewFile,

		OpenKey: Open,

		OpenFatKey: OpenFat,
	}
}

// NewFatFile creates a new FatFile for accessing all the Mach-O images in a
// universal binary. The Mach-O binary is expected to start at position 0 in
// the ReaderAt.
var NewFatFile = macho.NewFatFile

// NewFile creates a new File for accessing a Mach-O binary in an underlying reader.
// The Mach-O binary is expected to start at position 0 in the ReaderAt.
var NewFile = macho.NewFile

// Open opens the named file using os.Open and prepares it for use as a Mach-O binary.
var Open = macho.Open

// OpenFat opens the named file using os.Open and prepares it for use as a Mach-O
// universal binary.
var OpenFat = macho.OpenFat
