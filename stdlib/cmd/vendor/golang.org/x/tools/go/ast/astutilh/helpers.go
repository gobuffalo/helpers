package astutilh

import (
	"cmd/vendor/golang.org/x/tools/go/ast/astutil"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	AddImportKey = "AddImport"

	AddNamedImportKey = "AddNamedImport"

	ApplyKey = "Apply"

	DeleteImportKey = "DeleteImport"

	DeleteNamedImportKey = "DeleteNamedImport"

	ImportsKey = "Imports"

	NodeDescriptionKey = "NodeDescription"

	PathEnclosingIntervalKey = "PathEnclosingInterval"

	RewriteImportKey = "RewriteImport"

	UnparenKey = "Unparen"

	UsesImportKey = "UsesImport"
)

func New() hctx.Map {
	return hctx.Map{

		AddImportKey: AddImport,

		AddNamedImportKey: AddNamedImport,

		ApplyKey: Apply,

		DeleteImportKey: DeleteImport,

		DeleteNamedImportKey: DeleteNamedImport,

		ImportsKey: Imports,

		NodeDescriptionKey: NodeDescription,

		PathEnclosingIntervalKey: PathEnclosingInterval,

		RewriteImportKey: RewriteImport,

		UnparenKey: Unparen,

		UsesImportKey: UsesImport,
	}
}

// AddImport adds the import path to the file f, if absent.
var AddImport = astutil.AddImport

// AddNamedImport adds the import with the given name and path to the file f, if absent.
// If name is not empty, it is used to rename the import.
//
// For example, calling
// 	AddNamedImport(fset, f, &#34;pathpkg&#34;, &#34;path&#34;)
// adds
// 	import pathpkg &#34;path&#34;
var AddNamedImport = astutil.AddNamedImport

// Apply traverses a syntax tree recursively, starting with root,
// and calling pre and post for each node as described below.
// Apply returns the syntax tree, possibly modified.
//
// If pre is not nil, it is called for each node before the node&#39;s
// children are traversed (pre-order). If pre returns false, no
// children are traversed, and post is not called for that node.
//
// If post is not nil, and a prior call of pre didn&#39;t return false,
// post is called for each node after its children are traversed
// (post-order). If post returns false, traversal is terminated and
// Apply returns immediately.
//
// Only fields that refer to AST nodes are considered children;
// i.e., token.Pos, Scopes, Objects, and fields of basic types
// (strings, etc.) are ignored.
//
// Children are traversed in the order in which they appear in the
// respective node&#39;s struct definition. A package&#39;s files are
// traversed in the filenames&#39; alphabetical order.
var Apply = astutil.Apply

// DeleteImport deletes the import path from the file f, if present.
// If there are duplicate import declarations, all matching ones are deleted.
var DeleteImport = astutil.DeleteImport

// DeleteNamedImport deletes the import with the given name and path from the file f, if present.
// If there are duplicate import declarations, all matching ones are deleted.
var DeleteNamedImport = astutil.DeleteNamedImport

// Imports returns the file imports grouped by paragraph.
var Imports = astutil.Imports

// NodeDescription returns a description of the concrete type of n suitable
// for a user interface.
//
// TODO(adonovan): in some cases (e.g. Field, FieldList, Ident,
// StarExpr) we could be much more specific given the path to the AST
// root.  Perhaps we should do that.
var NodeDescription = astutil.NodeDescription

// PathEnclosingInterval returns the node that encloses the source
// interval [start, end), and all its ancestors up to the AST root.
//
// The definition of &#34;enclosing&#34; used by this function considers
// additional whitespace abutting a node to be enclosed by it.
// In this example:
//
//              z := x + y // add them
//                   &lt;-A-&gt;
//                  &lt;----B-----&gt;
//
// the ast.BinaryExpr(+) node is considered to enclose interval B
// even though its [Pos()..End()) is actually only interval A.
// This behaviour makes user interfaces more tolerant of imperfect
// input.
//
// This function treats tokens as nodes, though they are not included
// in the result. e.g. PathEnclosingInterval(&#34;+&#34;) returns the
// enclosing ast.BinaryExpr(&#34;x + y&#34;).
//
// If start==end, the 1-char interval following start is used instead.
//
// The &#39;exact&#39; result is true if the interval contains only path[0]
// and perhaps some adjacent whitespace.  It is false if the interval
// overlaps multiple children of path[0], or if it contains only
// interior whitespace of path[0].
// In this example:
//
//              z := x + y // add them
//                &lt;--C--&gt;     &lt;---E--&gt;
//                  ^
//                  D
//
// intervals C, D and E are inexact.  C is contained by the
// z-assignment statement, because it spans three of its children (:=,
// x, +).  So too is the 1-char interval D, because it contains only
// interior whitespace of the assignment.  E is considered interior
// whitespace of the BlockStmt containing the assignment.
//
// Precondition: [start, end) both lie within the same file as root.
// TODO(adonovan): return (nil, false) in this case and remove precond.
// Requires FileSet; see loader.tokenFileContainsPos.
//
// Postcondition: path is never nil; it always contains at least &#39;root&#39;.
var PathEnclosingInterval = astutil.PathEnclosingInterval

// RewriteImport rewrites any import of path oldPath to path newPath.
var RewriteImport = astutil.RewriteImport

// Unparen returns e with any enclosing parentheses stripped.
var Unparen = astutil.Unparen

// UsesImport reports whether a given import is used.
var UsesImport = astutil.UsesImport
