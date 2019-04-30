package quickh

import (
	"testing/quick"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	CheckKey = "Check"

	CheckEqualKey = "CheckEqual"

	ValueKey = "Value"
)

func New() hctx.Map {
	return hctx.Map{

		CheckKey: Check,

		CheckEqualKey: CheckEqual,

		ValueKey: Value,
	}
}

// Check looks for an input to f, any function that returns bool,
// such that f returns false. It calls f repeatedly, with arbitrary
// values for each argument. If f returns false on a given input,
// Check returns that input as a *CheckError.
// For example:
//
// 	func TestOddMultipleOfThree(t *testing.T) {
// 		f := func(x int) bool {
// 			y := OddMultipleOfThree(x)
// 			return y%2 == 1 &amp;&amp; y%3 == 0
// 		}
// 		if err := quick.Check(f, nil); err != nil {
// 			t.Error(err)
// 		}
// 	}
var Check = quick.Check

// CheckEqual looks for an input on which f and g return different results.
// It calls f and g repeatedly with arbitrary values for each argument.
// If f and g return different answers, CheckEqual returns a *CheckEqualError
// describing the input and the outputs.
var CheckEqual = quick.CheckEqual

// Value returns an arbitrary value of the given type.
// If the type implements the Generator interface, that will be used.
// Note: To create arbitrary values for structs, all the fields must be exported.
var Value = quick.Value
