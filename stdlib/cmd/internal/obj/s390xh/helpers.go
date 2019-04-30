package s390xh

import (
	"cmd/internal/obj/s390x"

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

var DRconv = s390x.DRconv
