package importerh

import (
	"go/importer"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DefaultKey = "Default"

	ForKey = "For"

	ForCompilerKey = "ForCompiler"
)

func New() hctx.Map {
	return hctx.Map{

		DefaultKey: Default,

		ForKey: For,

		ForCompilerKey: ForCompiler,
	}
}

// Default returns an Importer for the compiler that built the running binary.
// If available, the result implements types.ImporterFrom.
var Default = importer.Default

// For calls ForCompiler with a new FileSet.
//
// Deprecated: use ForCompiler, which populates a FileSet
// with the positions of objects created by the importer.
var For = importer.For

// ForCompiler returns an Importer for importing from installed packages
// for the compilers &#34;gc&#34; and &#34;gccgo&#34;, or for importing directly
// from the source if the compiler argument is &#34;source&#34;. In this
// latter case, importing may fail under circumstances where the
// exported API is not entirely defined in pure Go source code
// (if the package API depends on cgo-defined entities, the type
// checker won&#39;t have access to those).
//
// If lookup is nil, the default package lookup mechanism for the
// given compiler is used, and the resulting importer attempts
// to resolve relative and absolute import paths to canonical
// import path IDs before finding the imported file.
//
// If lookup is non-nil, then the returned importer calls lookup
// each time it needs to resolve an import path. In this mode
// the importer can only be invoked with canonical import paths
// (not relative or absolute ones); it is assumed that the translation
// to canonical import paths is being done by the client of the
// importer.
var ForCompiler = importer.ForCompiler
