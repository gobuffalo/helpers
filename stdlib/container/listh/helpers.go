package listh

import (
	"container/list"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewKey = "New"
)

func New() hctx.Map {
	return hctx.Map{

		NewKey: New,
	}
}

// New returns an initialized list.
var New = list.New
