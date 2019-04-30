package mipsh

import (
	"cmd/compile/internal/mips"

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

var Init = mips.Init
