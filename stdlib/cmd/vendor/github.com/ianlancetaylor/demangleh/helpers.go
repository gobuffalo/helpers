package demangleh

import (
	"cmd/vendor/github.com/ianlancetaylor/demangle"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	ASTToStringKey = "ASTToString"

	FilterKey = "Filter"

	ToASTKey = "ToAST"

	ToStringKey = "ToString"
)

func New() hctx.Map {
	return hctx.Map{

		ASTToStringKey: ASTToString,

		FilterKey: Filter,

		ToASTKey: ToAST,

		ToStringKey: ToString,
	}
}

// ASTToString returns the demangled name of the AST.
var ASTToString = demangle.ASTToString

// Filter demangles a C++ symbol name, returning the human-readable C++ name.
// If any error occurs during demangling, the input string is returned.
var Filter = demangle.Filter

// ToAST demangles a C++ symbol name into an abstract syntax tree
// representing the symbol.
// If the NoParams option is passed, and the name has a function type,
// the parameter types are not demangled.
// If the name does not appear to be a C++ symbol name at all, the
// error will be ErrNotMangledName.
var ToAST = demangle.ToAST

// ToString demangles a C++ symbol name, returning human-readable C++
// name or an error.
// If the name does not appear to be a C++ symbol name at all, the
// error will be ErrNotMangledName.
var ToString = demangle.ToString
