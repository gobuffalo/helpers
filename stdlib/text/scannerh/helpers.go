package scannerh

import (
	"text/scanner"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	TokenStringKey = "TokenString"
)

func New() hctx.Map {
	return hctx.Map{

		TokenStringKey: TokenString,
	}
}

// TokenString returns a printable string for a token or Unicode character.
var TokenString = scanner.TokenString
