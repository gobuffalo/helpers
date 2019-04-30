package asth

import (
	"go/ast"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	FileExportsKey = "FileExports"

	FilterDeclKey = "FilterDecl"

	FilterFileKey = "FilterFile"

	FilterPackageKey = "FilterPackage"

	FprintKey = "Fprint"

	InspectKey = "Inspect"

	IsExportedKey = "IsExported"

	MergePackageFilesKey = "MergePackageFiles"

	NewCommentMapKey = "NewCommentMap"

	NewIdentKey = "NewIdent"

	NewObjKey = "NewObj"

	NewPackageKey = "NewPackage"

	NewScopeKey = "NewScope"

	NotNilFilterKey = "NotNilFilter"

	PackageExportsKey = "PackageExports"

	PrintKey = "Print"

	SortImportsKey = "SortImports"

	WalkKey = "Walk"
)

func New() hctx.Map {
	return hctx.Map{

		FileExportsKey: FileExports,

		FilterDeclKey: FilterDecl,

		FilterFileKey: FilterFile,

		FilterPackageKey: FilterPackage,

		FprintKey: Fprint,

		InspectKey: Inspect,

		IsExportedKey: IsExported,

		MergePackageFilesKey: MergePackageFiles,

		NewCommentMapKey: NewCommentMap,

		NewIdentKey: NewIdent,

		NewObjKey: NewObj,

		NewPackageKey: NewPackage,

		NewScopeKey: NewScope,

		NotNilFilterKey: NotNilFilter,

		PackageExportsKey: PackageExports,

		PrintKey: Print,

		SortImportsKey: SortImports,

		WalkKey: Walk,
	}
}

// FileExports trims the AST for a Go source file in place such that
// only exported nodes remain: all top-level identifiers which are not exported
// and their associated information (such as type, initial value, or function
// body) are removed. Non-exported fields and methods of exported types are
// stripped. The File.Comments list is not changed.
//
// FileExports reports whether there are exported declarations.
var FileExports = ast.FileExports

// FilterDecl trims the AST for a Go declaration in place by removing
// all names (including struct field and interface method names, but
// not from parameter lists) that don&#39;t pass through the filter f.
//
// FilterDecl reports whether there are any declared names left after
// filtering.
var FilterDecl = ast.FilterDecl

// FilterFile trims the AST for a Go file in place by removing all
// names from top-level declarations (including struct field and
// interface method names, but not from parameter lists) that don&#39;t
// pass through the filter f. If the declaration is empty afterwards,
// the declaration is removed from the AST. Import declarations are
// always removed. The File.Comments list is not changed.
//
// FilterFile reports whether there are any top-level declarations
// left after filtering.
var FilterFile = ast.FilterFile

// FilterPackage trims the AST for a Go package in place by removing
// all names from top-level declarations (including struct field and
// interface method names, but not from parameter lists) that don&#39;t
// pass through the filter f. If the declaration is empty afterwards,
// the declaration is removed from the AST. The pkg.Files list is not
// changed, so that file names and top-level package comments don&#39;t get
// lost.
//
// FilterPackage reports whether there are any top-level declarations
// left after filtering.
var FilterPackage = ast.FilterPackage

// Fprint prints the (sub-)tree starting at AST node x to w.
// If fset != nil, position information is interpreted relative
// to that file set. Otherwise positions are printed as integer
// values (file set specific offsets).
//
// A non-nil FieldFilter f may be provided to control the output:
// struct fields for which f(fieldname, fieldvalue) is true are
// printed; all others are filtered from the output. Unexported
// struct fields are never printed.
var Fprint = ast.Fprint

// Inspect traverses an AST in depth-first order: It starts by calling
// f(node); node must not be nil. If f returns true, Inspect invokes f
// recursively for each of the non-nil children of node, followed by a
// call of f(nil).
var Inspect = ast.Inspect

// IsExported reports whether name is an exported Go symbol
// (that is, whether it begins with an upper-case letter).
var IsExported = ast.IsExported

// MergePackageFiles creates a file AST by merging the ASTs of the
// files belonging to a package. The mode flags control merging behavior.
var MergePackageFiles = ast.MergePackageFiles

// NewCommentMap creates a new comment map by associating comment groups
// of the comments list with the nodes of the AST specified by node.
//
// A comment group g is associated with a node n if:
//
//   - g starts on the same line as n ends
//   - g starts on the line immediately following n, and there is
//     at least one empty line after g and before the next node
//   - g starts before n and is not associated to the node before n
//     via the previous rules
//
// NewCommentMap tries to associate a comment group to the &#34;largest&#34;
// node possible: For instance, if the comment is a line comment
// trailing an assignment, the comment is associated with the entire
// assignment rather than just the last operand in the assignment.
var NewCommentMap = ast.NewCommentMap

// NewIdent creates a new Ident without position.
// Useful for ASTs generated by code other than the Go parser.
var NewIdent = ast.NewIdent

// NewObj creates a new object of a given kind and name.
var NewObj = ast.NewObj

// NewPackage creates a new Package node from a set of File nodes. It resolves
// unresolved identifiers across files and updates each file&#39;s Unresolved list
// accordingly. If a non-nil importer and universe scope are provided, they are
// used to resolve identifiers not declared in any of the package files. Any
// remaining unresolved identifiers are reported as undeclared. If the files
// belong to different packages, one package name is selected and files with
// different package names are reported and then ignored.
// The result is a package node and a scanner.ErrorList if there were errors.
var NewPackage = ast.NewPackage

// NewScope creates a new scope nested in the outer scope.
var NewScope = ast.NewScope

// NotNilFilter returns true for field values that are not nil;
// it returns false otherwise.
var NotNilFilter = ast.NotNilFilter

// PackageExports trims the AST for a Go package in place such that
// only exported nodes remain. The pkg.Files list is not changed, so that
// file names and top-level package comments don&#39;t get lost.
//
// PackageExports reports whether there are exported declarations;
// it returns false otherwise.
var PackageExports = ast.PackageExports

// Print prints x to standard output, skipping nil fields.
// Print(fset, x) is the same as Fprint(os.Stdout, fset, x, NotNilFilter).
var Print = ast.Print

// SortImports sorts runs of consecutive import lines in import blocks in f.
// It also removes duplicate imports when it is possible to do so without data loss.
var SortImports = ast.SortImports

// Walk traverses an AST in depth-first order: It starts by calling
// v.Visit(node); node must not be nil. If the visitor w returned by
// v.Visit(node) is not nil, Walk is invoked recursively with visitor
// w for each of the non-nil children of node, followed by a call of
// w.Visit(nil).
var Walk = ast.Walk
