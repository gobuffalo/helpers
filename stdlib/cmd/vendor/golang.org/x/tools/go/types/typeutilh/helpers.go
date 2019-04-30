package typeutilh

import (
	"cmd/vendor/golang.org/x/tools/go/types/typeutil"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	CalleeKey = "Callee"

	DependenciesKey = "Dependencies"

	IntuitiveMethodSetKey = "IntuitiveMethodSet"

	MakeHasherKey = "MakeHasher"

	StaticCalleeKey = "StaticCallee"
)

func New() hctx.Map {
	return hctx.Map{

		CalleeKey: Callee,

		DependenciesKey: Dependencies,

		IntuitiveMethodSetKey: IntuitiveMethodSet,

		MakeHasherKey: MakeHasher,

		StaticCalleeKey: StaticCallee,
	}
}

// Callee returns the named target of a function call, if any:
// a function, method, builtin, or variable.
var Callee = typeutil.Callee

// Dependencies returns all dependencies of the specified packages.
//
// Dependent packages appear in topological order: if package P imports
// package Q, Q appears earlier than P in the result.
// The algorithm follows import statements in the order they
// appear in the source code, so the result is a total order.
var Dependencies = typeutil.Dependencies

// IntuitiveMethodSet returns the intuitive method set of a type T,
// which is the set of methods you can call on an addressable value of
// that type.
//
// The result always contains MethodSet(T), and is exactly MethodSet(T)
// for interface types and for pointer-to-concrete types.
// For all other concrete types T, the result additionally
// contains each method belonging to *T if there is no identically
// named method on T itself.
//
// This corresponds to user intuition about method sets;
// this function is intended only for user interfaces.
//
// The order of the result is as for types.MethodSet(T).
var IntuitiveMethodSet = typeutil.IntuitiveMethodSet

// MakeHasher returns a new Hasher instance.
var MakeHasher = typeutil.MakeHasher

// StaticCallee returns the target (function or method) of a static
// function call, if any. It returns nil for calls to builtins.
var StaticCallee = typeutil.StaticCallee
