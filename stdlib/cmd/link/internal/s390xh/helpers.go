package s390xh

import (
	"cmd/link/internal/s390x"

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

var Init = s390x.Init
