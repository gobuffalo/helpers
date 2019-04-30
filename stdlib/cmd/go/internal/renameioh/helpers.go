package renameioh

import (
	"cmd/go/internal/renameio"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	PatternKey = "Pattern"

	WriteFileKey = "WriteFile"

	WriteToFileKey = "WriteToFile"
)

func New() hctx.Map {
	return hctx.Map{

		PatternKey: Pattern,

		WriteFileKey: WriteFile,

		WriteToFileKey: WriteToFile,
	}
}

// Pattern returns a glob pattern that matches the unrenamed temporary files
// created when writing to filename.
var Pattern = renameio.Pattern

// WriteFile is like ioutil.WriteFile, but first writes data to an arbitrary
// file in the same directory as filename, then renames it atomically to the
// final name.
//
// That ensures that the final location, if it exists, is always a complete file.
var WriteFile = renameio.WriteFile

// WriteToFile is a variant of WriteFile that accepts the data as an io.Reader
// instead of a slice.
var WriteToFile = renameio.WriteToFile
