package gosymh

import (
	"debug/gosym"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewLineTableKey = "NewLineTable"

	NewTableKey = "NewTable"
)

func New() hctx.Map {
	return hctx.Map{

		NewLineTableKey: NewLineTable,

		NewTableKey: NewTable,
	}
}

// NewLineTable returns a new PC/line table
// corresponding to the encoded data.
// Text must be the start address of the
// corresponding text segment.
var NewLineTable = gosym.NewLineTable

// NewTable decodes the Go symbol table (the &#34;.gosymtab&#34; section in ELF),
// returning an in-memory representation.
// Starting with Go 1.3, the Go symbol table no longer includes symbol data.
var NewTable = gosym.NewTable
