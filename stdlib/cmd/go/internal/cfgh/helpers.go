package cfgh

import (
	"cmd/go/internal/cfg"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DefaultCCKey = "DefaultCC"

	DefaultCXXKey = "DefaultCXX"
)

func New() hctx.Map {
	return hctx.Map{

		DefaultCCKey: DefaultCC,

		DefaultCXXKey: DefaultCXX,
	}
}

var DefaultCC = cfg.DefaultCC

var DefaultCXX = cfg.DefaultCXX
