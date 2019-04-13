package env

import (
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/helpers/hctx"
)

func New() hctx.Map {
	return hctx.Map{
		"env":   Env,
		"envOr": EnvOr,
	}
}

var Env = envy.MustGet
var EnvOr = envy.Get
