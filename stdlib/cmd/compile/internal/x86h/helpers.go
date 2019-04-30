package x86h

import (
	"cmd/compile/internal/x86"

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

var Init = x86.Init
