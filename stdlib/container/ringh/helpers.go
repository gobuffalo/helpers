package ringh

import (
	"container/ring"

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

// New creates a ring of n elements.
var New = ring.New
