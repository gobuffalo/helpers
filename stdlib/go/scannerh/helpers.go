package scannerh

import (
	"go/scanner"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	PrintErrorKey = "PrintError"
)

func New() hctx.Map {
	return hctx.Map{

		PrintErrorKey: PrintError,
	}
}

// PrintError is a utility function that prints a list of errors to w,
// one error per line, if the err parameter is an ErrorList. Otherwise
// it prints the err string.
var PrintError = scanner.PrintError
