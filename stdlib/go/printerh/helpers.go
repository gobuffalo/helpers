package printerh

import (
	"go/printer"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	FprintKey = "Fprint"
)

func New() hctx.Map {
	return hctx.Map{

		FprintKey: Fprint,
	}
}

// Fprint &#34;pretty-prints&#34; an AST node to output.
// It calls Config.Fprint with default settings.
// Note that gofmt uses tabs for indentation but spaces for alignment;
// use format.Node (package go/format) for output that matches gofmt.
var Fprint = printer.Fprint
