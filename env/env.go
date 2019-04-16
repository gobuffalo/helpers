package env

import (
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	EnvKey   = "env"
	EnvOrKey = "envOr"
)

func New() hctx.Map {
	return hctx.Map{
		EnvKey:   Env,
		EnvOrKey: EnvOr,
	}
}

var Env = envy.MustGet
var EnvOr = envy.Get
