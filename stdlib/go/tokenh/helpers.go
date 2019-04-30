package tokenh

import (
	"go/token"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	LookupKey = "Lookup"

	NewFileSetKey = "NewFileSet"
)

func New() hctx.Map {
	return hctx.Map{

		LookupKey: Lookup,

		NewFileSetKey: NewFileSet,
	}
}

// Lookup maps an identifier to its keyword token or IDENT (if not a keyword).
var Lookup = token.Lookup

// NewFileSet creates a new file set.
var NewFileSet = token.NewFileSet
