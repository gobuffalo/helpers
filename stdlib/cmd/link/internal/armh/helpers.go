package armh

import (
	"cmd/link/internal/arm"

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

var Init = arm.Init
