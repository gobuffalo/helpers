package asmh

import (
	"cmd/asm/internal/asm"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewParserKey = "NewParser"
)

func New() hctx.Map {
	return hctx.Map{

		NewParserKey: NewParser,
	}
}

var NewParser = asm.NewParser
