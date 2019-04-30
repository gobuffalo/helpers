package unsafeh

import (
	"unsafe"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AlignofKey = "Alignof"

	OffsetofKey = "Offsetof"

	SizeofKey = "Sizeof"
)

func New() hctx.Map {
	return hctx.Map{

		AlignofKey: Alignof,

		OffsetofKey: Offsetof,

		SizeofKey: Sizeof,
	}
}

// Alignof takes an expression x of any type and returns the required alignment
// of a hypothetical variable v as if v was declared via var v = x.
// It is the largest value m such that the address of v is always zero mod m.
// It is the same as the value returned by reflect.TypeOf(x).Align().
// As a special case, if a variable s is of struct type and f is a field
// within that struct, then Alignof(s.f) will return the required alignment
// of a field of that type within a struct. This case is the same as the
// value returned by reflect.TypeOf(s.f).FieldAlign().
// The return value of Alignof is a Go constant.
var Alignof = unsafe.Alignof

// Offsetof returns the offset within the struct of the field represented by x,
// which must be of the form structValue.field. In other words, it returns the
// number of bytes between the start of the struct and the start of the field.
// The return value of Offsetof is a Go constant.
var Offsetof = unsafe.Offsetof

// Sizeof takes an expression x of any type and returns the size in bytes
// of a hypothetical variable v as if v was declared via var v = x.
// The size does not include any memory possibly referenced by x.
// For instance, if x is a slice, Sizeof returns the size of the slice
// descriptor, not the size of the memory referenced by the slice.
// The return value of Sizeof is a Go constant.
var Sizeof = unsafe.Sizeof
