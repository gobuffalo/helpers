package parserh

import (
	"go/parser"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ParseDirKey = "ParseDir"

	ParseExprKey = "ParseExpr"

	ParseExprFromKey = "ParseExprFrom"

	ParseFileKey = "ParseFile"
)

func New() hctx.Map {
	return hctx.Map{

		ParseDirKey: ParseDir,

		ParseExprKey: ParseExpr,

		ParseExprFromKey: ParseExprFrom,

		ParseFileKey: ParseFile,
	}
}

// ParseDir calls ParseFile for all files with names ending in &#34;.go&#34; in the
// directory specified by path and returns a map of package name -&gt; package
// AST with all the packages found.
//
// If filter != nil, only the files with os.FileInfo entries passing through
// the filter (and ending in &#34;.go&#34;) are considered. The mode bits are passed
// to ParseFile unchanged. Position information is recorded in fset, which
// must not be nil.
//
// If the directory couldn&#39;t be read, a nil map and the respective error are
// returned. If a parse error occurred, a non-nil but incomplete map and the
// first error encountered are returned.
var ParseDir = parser.ParseDir

// ParseExpr is a convenience function for obtaining the AST of an expression x.
// The position information recorded in the AST is undefined. The filename used
// in error messages is the empty string.
var ParseExpr = parser.ParseExpr

// ParseExprFrom is a convenience function for parsing an expression.
// The arguments have the same meaning as for ParseFile, but the source must
// be a valid Go (type or value) expression. Specifically, fset must not
// be nil.
var ParseExprFrom = parser.ParseExprFrom

// ParseFile parses the source code of a single Go source file and returns
// the corresponding ast.File node. The source code may be provided via
// the filename of the source file, or via the src parameter.
//
// If src != nil, ParseFile parses the source from src and the filename is
// only used when recording position information. The type of the argument
// for the src parameter must be string, []byte, or io.Reader.
// If src == nil, ParseFile parses the file specified by filename.
//
// The mode parameter controls the amount of source text parsed and other
// optional parser functionality. Position information is recorded in the
// file set fset, which must not be nil.
//
// If the source couldn&#39;t be read, the returned AST is nil and the error
// indicates the specific failure. If the source was read but syntax
// errors were found, the result is a partial AST (with ast.Bad* nodes
// representing the fragments of erroneous source code). Multiple errors
// are returned via a scanner.ErrorList which is sorted by file position.
var ParseFile = parser.ParseFile
