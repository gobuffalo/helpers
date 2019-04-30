package synch

import (
	"sync"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewCondKey = "NewCond"
)

func New() hctx.Map {
	return hctx.Map{

		NewCondKey: NewCond,
	}
}

// NewCond returns a new Cond with Locker l.
var NewCond = sync.NewCond
