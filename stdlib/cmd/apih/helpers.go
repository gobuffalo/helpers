package apih

import (
	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewWalkerKey = "NewWalker"
)

func New() hctx.Map {
	return hctx.Map{

		NewWalkerKey: NewWalker,
	}
}

var NewWalker = api.NewWalker
