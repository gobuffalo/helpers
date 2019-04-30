package objectpathh

import (
	"cmd/vendor/golang.org/x/tools/go/types/objectpath"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	ForKey = "For"

	ObjectKey = "Object"
)

func New() hctx.Map {
	return hctx.Map{

		ForKey: For,

		ObjectKey: Object,
	}
}

// The For function returns the path to an object relative to its package,
// or an error if the object is not accessible from the package&#39;s Scope.
//
// The For function guarantees to return a path only for the following objects:
// - package-level types
// - exported package-level non-types
// - methods
// - parameter and result variables
// - struct fields
// These objects are sufficient to define the API of their package.
// The objects described by a package&#39;s export data are drawn from this set.
//
// For does not return a path for predeclared names, imported package
// names, local names, and unexported package-level names (except
// types).
//
// Example: given this definition,
//
// 	package p
//
// 	type T interface {
// 		f() (a string, b struct{ X int })
// 	}
//
// For(X) would return a path that denotes the following sequence of operations:
//
//    p.Scope().Lookup(&#34;T&#34;)				(TypeName T)
//    .Type().Underlying().Method(0).			(method Func f)
//    .Type().Results().At(1)				(field Var b)
//    .Type().Field(0)					(field Var X)
//
// where p is the package (*types.Package) to which X belongs.
var For = objectpath.For

// Object returns the object denoted by path p within the package pkg.
var Object = objectpath.Object
