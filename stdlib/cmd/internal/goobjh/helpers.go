package goobjh

import (
	"cmd/internal/goobj"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ParseKey = "Parse"
)

func New() hctx.Map {
	return hctx.Map{

		ParseKey: Parse,
	}
}

// Parse parses an object file or archive from f,
// assuming that its import path is pkgpath.
var Parse = goobj.Parse
