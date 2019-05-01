package reflecth

import (
	"reflect"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AppendKey = "Append"

	AppendSliceKey = "AppendSlice"

	ArrayOfKey = "ArrayOf"

	ChanOfKey = "ChanOf"

	CopyKey = "Copy"

	DeepEqualKey = "DeepEqual"

	FuncOfKey = "FuncOf"

	IndirectKey = "Indirect"

	MakeChanKey = "MakeChan"

	MakeFuncKey = "MakeFunc"

	MakeMapKey = "MakeMap"

	MakeMapWithSizeKey = "MakeMapWithSize"

	MakeSliceKey = "MakeSlice"

	MapOfKey = "MapOf"

	NewAtKey = "NewAt"

	PtrToKey = "PtrTo"

	SelectKey = "Select"

	SliceOfKey = "SliceOf"

	StructOfKey = "StructOf"

	SwapperKey = "Swapper"

	TypeOfKey = "TypeOf"

	ValueOfKey = "ValueOf"

	ZeroKey = "Zero"
)

func New() hctx.Map {
	return hctx.Map{

		AppendKey: Append,

		AppendSliceKey: AppendSlice,

		ArrayOfKey: ArrayOf,

		ChanOfKey: ChanOf,

		CopyKey: Copy,

		DeepEqualKey: DeepEqual,

		FuncOfKey: FuncOf,

		IndirectKey: Indirect,

		MakeChanKey: MakeChan,

		MakeFuncKey: MakeFunc,

		MakeMapKey: MakeMap,

		MakeMapWithSizeKey: MakeMapWithSize,

		MakeSliceKey: MakeSlice,

		MapOfKey: MapOf,

		NewAtKey: NewAt,

		PtrToKey: PtrTo,

		SelectKey: Select,

		SliceOfKey: SliceOf,

		StructOfKey: StructOf,

		SwapperKey: Swapper,

		TypeOfKey: TypeOf,

		ValueOfKey: ValueOf,

		ZeroKey: Zero,
	}
}

// Append appends the values x to a slice s and returns the resulting slice.
// As in Go, each x&#39;s value must be assignable to the slice&#39;s element type.
var Append = reflect.Append

// AppendSlice appends a slice t to a slice s and returns the resulting slice.
// The slices s and t must have the same element type.
var AppendSlice = reflect.AppendSlice

// ArrayOf returns the array type with the given count and element type.
// For example, if t represents int, ArrayOf(5, t) represents [5]int.
//
// If the resulting type would be larger than the available address space,
// ArrayOf panics.
var ArrayOf = reflect.ArrayOf

// ChanOf returns the channel type with the given direction and element type.
// For example, if t represents int, ChanOf(RecvDir, t) represents &lt;-chan int.
//
// The gc runtime imposes a limit of 64 kB on channel element types.
// If t&#39;s size is equal to or exceeds this limit, ChanOf panics.
var ChanOf = reflect.ChanOf

// Copy copies the contents of src into dst until either
// dst has been filled or src has been exhausted.
// It returns the number of elements copied.
// Dst and src each must have kind Slice or Array, and
// dst and src must have the same element type.
//
// As a special case, src can have kind String if the element type of dst is kind Uint8.
var Copy = reflect.Copy

// DeepEqual reports whether x and y are ``deeply equal,&#39;&#39; defined as follows.
// Two values of identical type are deeply equal if one of the following cases applies.
// Values of distinct types are never deeply equal.
//
// Array values are deeply equal when their corresponding elements are deeply equal.
//
// Struct values are deeply equal if their corresponding fields,
// both exported and unexported, are deeply equal.
//
// Func values are deeply equal if both are nil; otherwise they are not deeply equal.
//
// Interface values are deeply equal if they hold deeply equal concrete values.
//
// Map values are deeply equal when all of the following are true:
// they are both nil or both non-nil, they have the same length,
// and either they are the same map object or their corresponding keys
// (matched using Go equality) map to deeply equal values.
//
// Pointer values are deeply equal if they are equal using Go&#39;s == operator
// or if they point to deeply equal values.
//
// Slice values are deeply equal when all of the following are true:
// they are both nil or both non-nil, they have the same length,
// and either they point to the same initial entry of the same underlying array
// (that is, &amp;x[0] == &amp;y[0]) or their corresponding elements (up to length) are deeply equal.
// Note that a non-nil empty slice and a nil slice (for example, []byte{} and []byte(nil))
// are not deeply equal.
//
// Other values - numbers, bools, strings, and channels - are deeply equal
// if they are equal using Go&#39;s == operator.
//
// In general DeepEqual is a recursive relaxation of Go&#39;s == operator.
// However, this idea is impossible to implement without some inconsistency.
// Specifically, it is possible for a value to be unequal to itself,
// either because it is of func type (uncomparable in general)
// or because it is a floating-point NaN value (not equal to itself in floating-point comparison),
// or because it is an array, struct, or interface containing
// such a value.
// On the other hand, pointer values are always equal to themselves,
// even if they point at or contain such problematic values,
// because they compare equal using Go&#39;s == operator, and that
// is a sufficient condition to be deeply equal, regardless of content.
// DeepEqual has been defined so that the same short-cut applies
// to slices and maps: if x and y are the same slice or the same map,
// they are deeply equal regardless of content.
//
// As DeepEqual traverses the data values it may find a cycle. The
// second and subsequent times that DeepEqual compares two pointer
// values that have been compared before, it treats the values as
// equal rather than examining the values to which they point.
// This ensures that DeepEqual terminates.
var DeepEqual = reflect.DeepEqual

// FuncOf returns the function type with the given argument and result types.
// For example if k represents int and e represents string,
// FuncOf([]Type{k}, []Type{e}, false) represents func(int) string.
//
// The variadic argument controls whether the function is variadic. FuncOf
// panics if the in[len(in)-1] does not represent a slice and variadic is
// true.
var FuncOf = reflect.FuncOf

// Indirect returns the value that v points to.
// If v is a nil pointer, Indirect returns a zero Value.
// If v is not a pointer, Indirect returns v.
var Indirect = reflect.Indirect

// MakeChan creates a new channel with the specified type and buffer size.
var MakeChan = reflect.MakeChan

// MakeFunc returns a new function of the given Type
// that wraps the function fn. When called, that new function
// does the following:
//
// 	- converts its arguments to a slice of Values.
// 	- runs results := fn(args).
// 	- returns the results as a slice of Values, one per formal result.
//
// The implementation fn can assume that the argument Value slice
// has the number and type of arguments given by typ.
// If typ describes a variadic function, the final Value is itself
// a slice representing the variadic arguments, as in the
// body of a variadic function. The result Value slice returned by fn
// must have the number and type of results given by typ.
//
// The Value.Call method allows the caller to invoke a typed function
// in terms of Values; in contrast, MakeFunc allows the caller to implement
// a typed function in terms of Values.
//
// The Examples section of the documentation includes an illustration
// of how to use MakeFunc to build a swap function for different types.
var MakeFunc = reflect.MakeFunc

// MakeMap creates a new map with the specified type.
var MakeMap = reflect.MakeMap

// MakeMapWithSize creates a new map with the specified type
// and initial space for approximately n elements.
var MakeMapWithSize = reflect.MakeMapWithSize

// MakeSlice creates a new zero-initialized slice value
// for the specified slice type, length, and capacity.
var MakeSlice = reflect.MakeSlice

// MapOf returns the map type with the given key and element types.
// For example, if k represents int and e represents string,
// MapOf(k, e) represents map[int]string.
//
// If the key type is not a valid map key type (that is, if it does
// not implement Go&#39;s == operator), MapOf panics.
var MapOf = reflect.MapOf

// NewAt returns a Value representing a pointer to a value of the
// specified type, using p as that pointer.
var NewAt = reflect.NewAt

// PtrTo returns the pointer type with element t.
// For example, if t represents type Foo, PtrTo(t) represents *Foo.
var PtrTo = reflect.PtrTo

// Select executes a select operation described by the list of cases.
// Like the Go select statement, it blocks until at least one of the cases
// can proceed, makes a uniform pseudo-random choice,
// and then executes that case. It returns the index of the chosen case
// and, if that case was a receive operation, the value received and a
// boolean indicating whether the value corresponds to a send on the channel
// (as opposed to a zero value received because the channel is closed).
var Select = reflect.Select

// SliceOf returns the slice type with element type t.
// For example, if t represents int, SliceOf(t) represents []int.
var SliceOf = reflect.SliceOf

// StructOf returns the struct type containing fields.
// The Offset and Index fields are ignored and computed as they would be
// by the compiler.
//
// StructOf currently does not generate wrapper methods for embedded
// fields and panics if passed unexported StructFields.
// These limitations may be lifted in a future version.
var StructOf = reflect.StructOf

// Swapper returns a function that swaps the elements in the provided
// slice.
//
// Swapper panics if the provided interface is not a slice.
var Swapper = reflect.Swapper

// TypeOf returns the reflection Type that represents the dynamic type of i.
// If i is a nil interface value, TypeOf returns nil.
var TypeOf = reflect.TypeOf

// ValueOf returns a new Value initialized to the concrete value
// stored in the interface i. ValueOf(nil) returns the zero Value.
var ValueOf = reflect.ValueOf

// Zero returns a Value representing the zero value for the specified type.
// The result is different from the zero value of the Value struct,
// which represents no value at all.
// For example, Zero(TypeOf(42)) returns a Value with Kind Int and value 0.
// The returned value is neither addressable nor settable.
var Zero = reflect.Zero
