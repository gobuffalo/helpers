package formath

import (
	"go/format"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NodeKey = "Node"

	SourceKey = "Source"
)

func New() hctx.Map {
	return hctx.Map{

		NodeKey: Node,

		SourceKey: Source,
	}
}

// Node formats node in canonical gofmt style and writes the result to dst.
//
// The node type must be *ast.File, *printer.CommentedNode, []ast.Decl,
// []ast.Stmt, or assignment-compatible to ast.Expr, ast.Decl, ast.Spec,
// or ast.Stmt. Node does not modify node. Imports are not sorted for
// nodes representing partial source files (for instance, if the node is
// not an *ast.File or a *printer.CommentedNode not wrapping an *ast.File).
//
// The function may return early (before the entire result is written)
// and return a formatting error, for instance due to an incorrect AST.
var Node = format.Node

// Source formats src in canonical gofmt style and returns the result
// or an (I/O or syntax) error. src is expected to be a syntactically
// correct Go source file, or a list of Go declarations or statements.
//
// If src is a partial source file, the leading and trailing space of src
// is applied to the result (such that it has the same leading and trailing
// space as src), and the result is indented by the same amount as the first
// line of src containing code. Imports are not sorted for partial source files.
//
// Caution: Tools relying on consistent formatting based on the installed
// version of gofmt (for instance, such as for presubmit checks) should
// execute that gofmt binary instead of calling Source.
var Source = format.Source