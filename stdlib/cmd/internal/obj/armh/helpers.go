package armh

import (
	"cmd/internal/obj/arm"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DRconvKey = "DRconv"
)

func New() hctx.Map {
	return hctx.Map{

		DRconvKey: DRconv,
	}
}

var DRconv = arm.DRconv
