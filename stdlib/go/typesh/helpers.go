package typesh

import (
	"go/types"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AssertableToKey = "AssertableTo"

	AssignableToKey = "AssignableTo"

	ComparableKey = "Comparable"

	ConvertibleToKey = "ConvertibleTo"

	DefPredeclaredTestFuncsKey = "DefPredeclaredTestFuncs"

	DefaultKey = "Default"

	EvalKey = "Eval"

	ExprStringKey = "ExprString"

	IdKey = "Id"

	IdenticalKey = "Identical"

	IdenticalIgnoreTagsKey = "IdenticalIgnoreTags"

	ImplementsKey = "Implements"

	IsInterfaceKey = "IsInterface"

	LookupFieldOrMethodKey = "LookupFieldOrMethod"

	MissingMethodKey = "MissingMethod"

	NewArrayKey = "NewArray"

	NewChanKey = "NewChan"

	NewCheckerKey = "NewChecker"

	NewConstKey = "NewConst"

	NewFieldKey = "NewField"

	NewFuncKey = "NewFunc"

	NewInterfaceKey = "NewInterface"

	NewLabelKey = "NewLabel"

	NewMapKey = "NewMap"

	NewMethodSetKey = "NewMethodSet"

	NewNamedKey = "NewNamed"

	NewPackageKey = "NewPackage"

	NewParamKey = "NewParam"

	NewPkgNameKey = "NewPkgName"

	NewPointerKey = "NewPointer"

	NewScopeKey = "NewScope"

	NewSignatureKey = "NewSignature"

	NewSliceKey = "NewSlice"

	NewStructKey = "NewStruct"

	NewTupleKey = "NewTuple"

	NewTypeNameKey = "NewTypeName"

	NewVarKey = "NewVar"

	ObjectStringKey = "ObjectString"

	RelativeToKey = "RelativeTo"

	SelectionStringKey = "SelectionString"

	SizesForKey = "SizesFor"

	TypeStringKey = "TypeString"

	WriteExprKey = "WriteExpr"

	WriteSignatureKey = "WriteSignature"

	WriteTypeKey = "WriteType"
)

func New() hctx.Map {
	return hctx.Map{

		AssertableToKey: AssertableTo,

		AssignableToKey: AssignableTo,

		ComparableKey: Comparable,

		ConvertibleToKey: ConvertibleTo,

		DefPredeclaredTestFuncsKey: DefPredeclaredTestFuncs,

		DefaultKey: Default,

		EvalKey: Eval,

		ExprStringKey: ExprString,

		IdKey: Id,

		IdenticalKey: Identical,

		IdenticalIgnoreTagsKey: IdenticalIgnoreTags,

		ImplementsKey: Implements,

		IsInterfaceKey: IsInterface,

		LookupFieldOrMethodKey: LookupFieldOrMethod,

		MissingMethodKey: MissingMethod,

		NewArrayKey: NewArray,

		NewChanKey: NewChan,

		NewCheckerKey: NewChecker,

		NewConstKey: NewConst,

		NewFieldKey: NewField,

		NewFuncKey: NewFunc,

		NewInterfaceKey: NewInterface,

		NewLabelKey: NewLabel,

		NewMapKey: NewMap,

		NewMethodSetKey: NewMethodSet,

		NewNamedKey: NewNamed,

		NewPackageKey: NewPackage,

		NewParamKey: NewParam,

		NewPkgNameKey: NewPkgName,

		NewPointerKey: NewPointer,

		NewScopeKey: NewScope,

		NewSignatureKey: NewSignature,

		NewSliceKey: NewSlice,

		NewStructKey: NewStruct,

		NewTupleKey: NewTuple,

		NewTypeNameKey: NewTypeName,

		NewVarKey: NewVar,

		ObjectStringKey: ObjectString,

		RelativeToKey: RelativeTo,

		SelectionStringKey: SelectionString,

		SizesForKey: SizesFor,

		TypeStringKey: TypeString,

		WriteExprKey: WriteExpr,

		WriteSignatureKey: WriteSignature,

		WriteTypeKey: WriteType,
	}
}

// AssertableTo reports whether a value of type V can be asserted to have type T.
var AssertableTo = types.AssertableTo

// AssignableTo reports whether a value of type V is assignable to a variable of type T.
var AssignableTo = types.AssignableTo

// Comparable reports whether values of type T are comparable.
var Comparable = types.Comparable

// ConvertibleTo reports whether a value of type V is convertible to a value of type T.
var ConvertibleTo = types.ConvertibleTo

// DefPredeclaredTestFuncs defines the assert and trace built-ins.
// These built-ins are intended for debugging and testing of this
// package only.
var DefPredeclaredTestFuncs = types.DefPredeclaredTestFuncs

// Default returns the default &#34;typed&#34; type for an &#34;untyped&#34; type;
// it returns the incoming type for all other types. The default type
// for untyped nil is untyped nil.
var Default = types.Default

// Eval returns the type and, if constant, the value for the
// expression expr, evaluated at position pos of package pkg,
// which must have been derived from type-checking an AST with
// complete position information relative to the provided file
// set.
//
// If the expression contains function literals, their bodies
// are ignored (i.e., the bodies are not type-checked).
//
// If pkg == nil, the Universe scope is used and the provided
// position pos is ignored. If pkg != nil, and pos is invalid,
// the package scope is used. Otherwise, pos must belong to the
// package.
//
// An error is returned if pos is not within the package or
// if the node cannot be evaluated.
//
// Note: Eval should not be used instead of running Check to compute
// types and values, but in addition to Check. Eval will re-evaluate
// its argument each time, and it also does not know about the context
// in which an expression is used (e.g., an assignment). Thus, top-
// level untyped constants will return an untyped type rather then the
// respective context-specific type.
var Eval = types.Eval

// ExprString returns the (possibly shortened) string representation for x.
// Shortened representations are suitable for user interfaces but may not
// necessarily follow Go syntax.
var ExprString = types.ExprString

// Id returns name if it is exported, otherwise it
// returns the name qualified with the package path.
var Id = types.Id

// Identical reports whether x and y are identical types.
// Receivers of Signature types are ignored.
var Identical = types.Identical

// IdenticalIgnoreTags reports whether x and y are identical types if tags are ignored.
// Receivers of Signature types are ignored.
var IdenticalIgnoreTags = types.IdenticalIgnoreTags

// Implements reports whether type V implements interface T.
var Implements = types.Implements

// IsInterface reports whether typ is an interface type.
var IsInterface = types.IsInterface

// LookupFieldOrMethod looks up a field or method with given package and name
// in T and returns the corresponding *Var or *Func, an index sequence, and a
// bool indicating if there were any pointer indirections on the path to the
// field or method. If addressable is set, T is the type of an addressable
// variable (only matters for method lookups).
//
// The last index entry is the field or method index in the (possibly embedded)
// type where the entry was found, either:
//
// 	1) the list of declared methods of a named type; or
// 	2) the list of all methods (method set) of an interface type; or
// 	3) the list of fields of a struct type.
//
// The earlier index entries are the indices of the anonymous struct fields
// traversed to get to the found entry, starting at depth 0.
//
// If no entry is found, a nil object is returned. In this case, the returned
// index and indirect values have the following meaning:
//
// 	- If index != nil, the index sequence points to an ambiguous entry
// 	(the same name appeared more than once at the same embedding level).
//
// 	- If indirect is set, a method with a pointer receiver type was found
//      but there was no pointer on the path from the actual receiver type to
// 	the method&#39;s formal receiver base type, nor was the receiver addressable.
var LookupFieldOrMethod = types.LookupFieldOrMethod

// MissingMethod returns (nil, false) if V implements T, otherwise it
// returns a missing method required by T and whether it is missing or
// just has the wrong type.
//
// For non-interface types V, or if static is set, V implements T if all
// methods of T are present in V. Otherwise (V is an interface and static
// is not set), MissingMethod only checks that methods of T which are also
// present in V have matching types (e.g., for a type assertion x.(T) where
// x is of interface type V).
var MissingMethod = types.MissingMethod

// NewArray returns a new array type for the given element type and length.
var NewArray = types.NewArray

// NewChan returns a new channel type for the given direction and element type.
var NewChan = types.NewChan

// NewChecker returns a new Checker instance for a given package.
// Package files may be added incrementally via checker.Files.
var NewChecker = types.NewChecker

// NewConst returns a new constant with value val.
// The remaining arguments set the attributes found with all Objects.
var NewConst = types.NewConst

// NewField returns a new variable representing a struct field.
// For anonymous (embedded) fields, the name is the unqualified
// type name under which the field is accessible.
var NewField = types.NewField

// NewFunc returns a new function with the given signature, representing
// the function&#39;s type.
var NewFunc = types.NewFunc

// NewInterface returns a new (incomplete) interface for the given methods and embedded types.
// To compute the method set of the interface, Complete must be called.
var NewInterface = types.NewInterface

// NewLabel returns a new label.
var NewLabel = types.NewLabel

// NewMap returns a new map for the given key and element types.
var NewMap = types.NewMap

// NewMethodSet returns the method set for the given type T.
// It always returns a non-nil method set, even if it is empty.
var NewMethodSet = types.NewMethodSet

// NewNamed returns a new named type for the given type name, underlying type, and associated methods.
// If the given type name obj doesn&#39;t have a type yet, its type is set to the returned named type.
// The underlying type must not be a *Named.
var NewNamed = types.NewNamed

// NewPackage returns a new Package for the given package path and name.
// The package is not complete and contains no explicit imports.
var NewPackage = types.NewPackage

// NewParam returns a new variable representing a function parameter.
var NewParam = types.NewParam

// NewPkgName returns a new PkgName object representing an imported package.
// The remaining arguments set the attributes found with all Objects.
var NewPkgName = types.NewPkgName

// NewPointer returns a new pointer type for the given element (base) type.
var NewPointer = types.NewPointer

// NewScope returns a new, empty scope contained in the given parent
// scope, if any. The comment is for debugging only.
var NewScope = types.NewScope

// NewSignature returns a new function type for the given receiver, parameters,
// and results, either of which may be nil. If variadic is set, the function
// is variadic, it must have at least one parameter, and the last parameter
// must be of unnamed slice type.
var NewSignature = types.NewSignature

// NewSlice returns a new slice type for the given element type.
var NewSlice = types.NewSlice

// NewStruct returns a new struct with the given fields and corresponding field tags.
// If a field with index i has a tag, tags[i] must be that tag, but len(tags) may be
// only as long as required to hold the tag with the largest index i. Consequently,
// if no field has a tag, tags may be nil.
var NewStruct = types.NewStruct

// NewTuple returns a new tuple for the given variables.
var NewTuple = types.NewTuple

// NewTypeName returns a new type name denoting the given typ.
// The remaining arguments set the attributes found with all Objects.
//
// The typ argument may be a defined (Named) type or an alias type.
// It may also be nil such that the returned TypeName can be used as
// argument for NewNamed, which will set the TypeName&#39;s type as a side-
// effect.
var NewTypeName = types.NewTypeName

// NewVar returns a new variable.
// The arguments set the attributes found with all Objects.
var NewVar = types.NewVar

// ObjectString returns the string form of obj.
// The Qualifier controls the printing of
// package-level objects, and may be nil.
var ObjectString = types.ObjectString

// RelativeTo(pkg) returns a Qualifier that fully qualifies members of
// all packages other than pkg.
var RelativeTo = types.RelativeTo

// SelectionString returns the string form of s.
// The Qualifier controls the printing of
// package-level objects, and may be nil.
//
// Examples:
// 	&#34;field (T) f int&#34;
// 	&#34;method (T) f(X) Y&#34;
// 	&#34;method expr (T) f(X) Y&#34;
var SelectionString = types.SelectionString

// SizesFor returns the Sizes used by a compiler for an architecture.
// The result is nil if a compiler/architecture pair is not known.
//
// Supported architectures for compiler &#34;gc&#34;:
// &#34;386&#34;, &#34;arm&#34;, &#34;arm64&#34;, &#34;amd64&#34;, &#34;amd64p32&#34;, &#34;mips&#34;, &#34;mipsle&#34;,
// &#34;mips64&#34;, &#34;mips64le&#34;, &#34;ppc64&#34;, &#34;ppc64le&#34;, &#34;s390x&#34;.
var SizesFor = types.SizesFor

// TypeString returns the string representation of typ.
// The Qualifier controls the printing of
// package-level objects, and may be nil.
var TypeString = types.TypeString

// WriteExpr writes the (possibly shortened) string representation for x to buf.
// Shortened representations are suitable for user interfaces but may not
// necessarily follow Go syntax.
var WriteExpr = types.WriteExpr

// WriteSignature writes the representation of the signature sig to buf,
// without a leading &#34;func&#34; keyword.
// The Qualifier controls the printing of
// package-level objects, and may be nil.
var WriteSignature = types.WriteSignature

// WriteType writes the string representation of typ to buf.
// The Qualifier controls the printing of
// package-level objects, and may be nil.
var WriteType = types.WriteType
