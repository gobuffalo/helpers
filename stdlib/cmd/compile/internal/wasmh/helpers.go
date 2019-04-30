package wasmh

import (
	"cmd/compile/internal/wasm"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	InitKey = "Init"
)

func New() hctx.Map {
	return hctx.Map{

		InitKey: Init,
	}
}

var Init = wasm.Init
