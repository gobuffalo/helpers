package constanth

import (
	"go/constant"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	BinaryOpKey = "BinaryOp"

	BitLenKey = "BitLen"

	BoolValKey = "BoolVal"

	BytesKey = "Bytes"

	CompareKey = "Compare"

	DenomKey = "Denom"

	Float32ValKey = "Float32Val"

	Float64ValKey = "Float64Val"

	ImagKey = "Imag"

	Int64ValKey = "Int64Val"

	MakeBoolKey = "MakeBool"

	MakeFloat64Key = "MakeFloat64"

	MakeFromBytesKey = "MakeFromBytes"

	MakeFromLiteralKey = "MakeFromLiteral"

	MakeImagKey = "MakeImag"

	MakeInt64Key = "MakeInt64"

	MakeStringKey = "MakeString"

	MakeUint64Key = "MakeUint64"

	MakeUnknownKey = "MakeUnknown"

	NumKey = "Num"

	RealKey = "Real"

	ShiftKey = "Shift"

	SignKey = "Sign"

	StringValKey = "StringVal"

	ToComplexKey = "ToComplex"

	ToFloatKey = "ToFloat"

	ToIntKey = "ToInt"

	Uint64ValKey = "Uint64Val"

	UnaryOpKey = "UnaryOp"
)

func New() hctx.Map {
	return hctx.Map{

		BinaryOpKey: BinaryOp,

		BitLenKey: BitLen,

		BoolValKey: BoolVal,

		BytesKey: Bytes,

		CompareKey: Compare,

		DenomKey: Denom,

		Float32ValKey: Float32Val,

		Float64ValKey: Float64Val,

		ImagKey: Imag,

		Int64ValKey: Int64Val,

		MakeBoolKey: MakeBool,

		MakeFloat64Key: MakeFloat64,

		MakeFromBytesKey: MakeFromBytes,

		MakeFromLiteralKey: MakeFromLiteral,

		MakeImagKey: MakeImag,

		MakeInt64Key: MakeInt64,

		MakeStringKey: MakeString,

		MakeUint64Key: MakeUint64,

		MakeUnknownKey: MakeUnknown,

		NumKey: Num,

		RealKey: Real,

		ShiftKey: Shift,

		SignKey: Sign,

		StringValKey: StringVal,

		ToComplexKey: ToComplex,

		ToFloatKey: ToFloat,

		ToIntKey: ToInt,

		Uint64ValKey: Uint64Val,

		UnaryOpKey: UnaryOp,
	}
}

// BinaryOp returns the result of the binary expression x op y.
// The operation must be defined for the operands. If one of the
// operands is Unknown, the result is Unknown.
// BinaryOp doesn&#39;t handle comparisons or shifts; use Compare
// or Shift instead.
//
// To force integer division of Int operands, use op == token.QUO_ASSIGN
// instead of token.QUO; the result is guaranteed to be Int in this case.
// Division by zero leads to a run-time panic.
var BinaryOp = constant.BinaryOp

// BitLen returns the number of bits required to represent
// the absolute value x in binary representation; x must be an Int or an Unknown.
// If x is Unknown, the result is 0.
var BitLen = constant.BitLen

// BoolVal returns the Go boolean value of x, which must be a Bool or an Unknown.
// If x is Unknown, the result is false.
var BoolVal = constant.BoolVal

// Bytes returns the bytes for the absolute value of x in little-
// endian binary representation; x must be an Int.
var Bytes = constant.Bytes

// Compare returns the result of the comparison x op y.
// The comparison must be defined for the operands.
// If one of the operands is Unknown, the result is
// false.
var Compare = constant.Compare

// Denom returns the denominator of x; x must be Int, Float, or Unknown.
// If x is Unknown, or if it is too large or small to represent as a
// fraction, the result is Unknown. Otherwise the result is an Int &gt;= 1.
var Denom = constant.Denom

// Float32Val is like Float64Val but for float32 instead of float64.
var Float32Val = constant.Float32Val

// Float64Val returns the nearest Go float64 value of x and whether the result is exact;
// x must be numeric or an Unknown, but not Complex. For values too small (too close to 0)
// to represent as float64, Float64Val silently underflows to 0. The result sign always
// matches the sign of x, even for 0.
// If x is Unknown, the result is (0, false).
var Float64Val = constant.Float64Val

// Imag returns the imaginary part of x, which must be a numeric or unknown value.
// If x is Unknown, the result is Unknown.
var Imag = constant.Imag

// Int64Val returns the Go int64 value of x and whether the result is exact;
// x must be an Int or an Unknown. If the result is not exact, its value is undefined.
// If x is Unknown, the result is (0, false).
var Int64Val = constant.Int64Val

// MakeBool returns the Bool value for b.
var MakeBool = constant.MakeBool

// MakeFloat64 returns the Float value for x.
// If x is not finite, the result is an Unknown.
var MakeFloat64 = constant.MakeFloat64

// MakeFromBytes returns the Int value given the bytes of its little-endian
// binary representation. An empty byte slice argument represents 0.
var MakeFromBytes = constant.MakeFromBytes

// MakeFromLiteral returns the corresponding integer, floating-point,
// imaginary, character, or string value for a Go literal string. The
// tok value must be one of token.INT, token.FLOAT, token.IMAG,
// token.CHAR, or token.STRING. The final argument must be zero.
// If the literal string syntax is invalid, the result is an Unknown.
var MakeFromLiteral = constant.MakeFromLiteral

// MakeImag returns the Complex value x*i;
// x must be Int, Float, or Unknown.
// If x is Unknown, the result is Unknown.
var MakeImag = constant.MakeImag

// MakeInt64 returns the Int value for x.
var MakeInt64 = constant.MakeInt64

// MakeString returns the String value for s.
var MakeString = constant.MakeString

// MakeUint64 returns the Int value for x.
var MakeUint64 = constant.MakeUint64

// MakeUnknown returns the Unknown value.
var MakeUnknown = constant.MakeUnknown

// Num returns the numerator of x; x must be Int, Float, or Unknown.
// If x is Unknown, or if it is too large or small to represent as a
// fraction, the result is Unknown. Otherwise the result is an Int
// with the same sign as x.
var Num = constant.Num

// Real returns the real part of x, which must be a numeric or unknown value.
// If x is Unknown, the result is Unknown.
var Real = constant.Real

// Shift returns the result of the shift expression x op s
// with op == token.SHL or token.SHR (&lt;&lt; or &gt;&gt;). x must be
// an Int or an Unknown. If x is Unknown, the result is x.
var Shift = constant.Shift

// Sign returns -1, 0, or 1 depending on whether x &lt; 0, x == 0, or x &gt; 0;
// x must be numeric or Unknown. For complex values x, the sign is 0 if x == 0,
// otherwise it is != 0. If x is Unknown, the result is 1.
var Sign = constant.Sign

// StringVal returns the Go string value of x, which must be a String or an Unknown.
// If x is Unknown, the result is &#34;&#34;.
var StringVal = constant.StringVal

// ToComplex converts x to a Complex value if x is representable as a Complex.
// Otherwise it returns an Unknown.
var ToComplex = constant.ToComplex

// ToFloat converts x to a Float value if x is representable as a Float.
// Otherwise it returns an Unknown.
var ToFloat = constant.ToFloat

// ToInt converts x to an Int value if x is representable as an Int.
// Otherwise it returns an Unknown.
var ToInt = constant.ToInt

// Uint64Val returns the Go uint64 value of x and whether the result is exact;
// x must be an Int or an Unknown. If the result is not exact, its value is undefined.
// If x is Unknown, the result is (0, false).
var Uint64Val = constant.Uint64Val

// UnaryOp returns the result of the unary expression op y.
// The operation must be defined for the operand.
// If prec &gt; 0 it specifies the ^ (xor) result size in bits.
// If y is Unknown, the result is Unknown.
var UnaryOp = constant.UnaryOp
