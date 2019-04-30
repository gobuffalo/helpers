package cfgh

import (
	"cmd/vendor/golang.org/x/tools/go/cfg"
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

// New returns a new control-flow graph for the specified function body,
// which must be non-nil.
//
// The CFG builder calls mayReturn to determine whether a given function
// call may return.  For example, calls to panic, os.Exit, and log.Fatal
// do not return, so the builder can remove infeasible graph edges
// following such calls.  The builder calls mayReturn only for a
// CallExpr beneath an ExprStmt.
var New = cfg.New
