package inspectorh

import (
	"cmd/vendor/golang.org/x/tools/go/ast/inspector"
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

// New returns an Inspector for the specified syntax trees.
var New = inspector.New
