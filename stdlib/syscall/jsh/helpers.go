package jsh

import (
	"syscall/js"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	FuncOfKey = "FuncOf"

	GlobalKey = "Global"

	NullKey = "Null"

	TypedArrayOfKey = "TypedArrayOf"

	UndefinedKey = "Undefined"

	ValueOfKey = "ValueOf"
)

func New() hctx.Map {
	return hctx.Map{

		FuncOfKey: FuncOf,

		GlobalKey: Global,

		NullKey: Null,

		TypedArrayOfKey: TypedArrayOf,

		UndefinedKey: Undefined,

		ValueOfKey: ValueOf,
	}
}

// FuncOf returns a wrapped function.
//
// Invoking the JavaScript function will synchronously call the Go function fn with the value of JavaScript&#39;s
// &#34;this&#34; keyword and the arguments of the invocation.
// The return value of the invocation is the result of the Go function mapped back to JavaScript according to ValueOf.
//
// A wrapped function triggered during a call from Go to JavaScript gets executed on the same goroutine.
// A wrapped function triggered by JavaScript&#39;s event loop gets executed on an extra goroutine.
// Blocking operations in the wrapped function will block the event loop.
// As a consequence, if one wrapped function blocks, other wrapped funcs will not be processed.
// A blocking function should therefore explicitly start a new goroutine.
//
// Func.Release must be called to free up resources when the function will not be used any more.
var FuncOf = js.FuncOf

// Global returns the JavaScript global object, usually &#34;window&#34; or &#34;global&#34;.
var Global = js.Global

// Null returns the JavaScript value &#34;null&#34;.
var Null = js.Null

// TypedArrayOf returns a JavaScript typed array backed by the slice&#39;s underlying array.
//
// The supported types are []int8, []int16, []int32, []uint8, []uint16, []uint32, []float32 and []float64.
// Passing an unsupported value causes a panic.
//
// TypedArray.Release must be called to free up resources when the typed array will not be used any more.
var TypedArrayOf = js.TypedArrayOf

// Undefined returns the JavaScript value &#34;undefined&#34;.
var Undefined = js.Undefined

// ValueOf returns x as a JavaScript value:
//
//  | Go                     | JavaScript             |
//  | ---------------------- | ---------------------- |
//  | js.Value               | [its value]            |
//  | js.TypedArray          | typed array            |
//  | js.Func                | function               |
//  | nil                    | null                   |
//  | bool                   | boolean                |
//  | integers and floats    | number                 |
//  | string                 | string                 |
//  | []interface{}          | new array              |
//  | map[string]interface{} | new object             |
//
// Panics if x is not one of the expected types.
var ValueOf = js.ValueOf
