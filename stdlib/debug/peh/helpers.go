package peh

import (
	"debug/pe"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewFileKey = "NewFile"

	OpenKey = "Open"
)

func New() hctx.Map {
	return hctx.Map{

		NewFileKey: NewFile,

		OpenKey: Open,
	}
}

// NewFile creates a new File for accessing a PE binary in an underlying reader.
var NewFile = pe.NewFile

// Open opens the named file using os.Open and prepares it for use as a PE binary.
var Open = pe.Open
