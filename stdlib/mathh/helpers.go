package mathh

import (
	"math"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AbsKey = "Abs"

	AcosKey = "Acos"

	AcoshKey = "Acosh"

	AsinKey = "Asin"

	AsinhKey = "Asinh"

	AtanKey = "Atan"

	Atan2Key = "Atan2"

	AtanhKey = "Atanh"

	CbrtKey = "Cbrt"

	CeilKey = "Ceil"

	CopysignKey = "Copysign"

	CosKey = "Cos"

	CoshKey = "Cosh"

	DimKey = "Dim"

	ErfKey = "Erf"

	ErfcKey = "Erfc"

	ErfcinvKey = "Erfcinv"

	ErfinvKey = "Erfinv"

	ExpKey = "Exp"

	Exp2Key = "Exp2"

	Expm1Key = "Expm1"

	Float32bitsKey = "Float32bits"

	Float32frombitsKey = "Float32frombits"

	Float64bitsKey = "Float64bits"

	Float64frombitsKey = "Float64frombits"

	FloorKey = "Floor"

	FrexpKey = "Frexp"

	GammaKey = "Gamma"

	HypotKey = "Hypot"

	IlogbKey = "Ilogb"

	InfKey = "Inf"

	IsInfKey = "IsInf"

	IsNaNKey = "IsNaN"

	J0Key = "J0"

	J1Key = "J1"

	JnKey = "Jn"

	LdexpKey = "Ldexp"

	LgammaKey = "Lgamma"

	LogKey = "Log"

	Log10Key = "Log10"

	Log1pKey = "Log1p"

	Log2Key = "Log2"

	LogbKey = "Logb"

	MaxKey = "Max"

	MinKey = "Min"

	ModKey = "Mod"

	ModfKey = "Modf"

	NaNKey = "NaN"

	NextafterKey = "Nextafter"

	Nextafter32Key = "Nextafter32"

	PowKey = "Pow"

	Pow10Key = "Pow10"

	RemainderKey = "Remainder"

	RoundKey = "Round"

	RoundToEvenKey = "RoundToEven"

	SignbitKey = "Signbit"

	SinKey = "Sin"

	SincosKey = "Sincos"

	SinhKey = "Sinh"

	SqrtKey = "Sqrt"

	TanKey = "Tan"

	TanhKey = "Tanh"

	TruncKey = "Trunc"

	Y0Key = "Y0"

	Y1Key = "Y1"

	YnKey = "Yn"
)

func New() hctx.Map {
	return hctx.Map{

		AbsKey: Abs,

		AcosKey: Acos,

		AcoshKey: Acosh,

		AsinKey: Asin,

		AsinhKey: Asinh,

		AtanKey: Atan,

		Atan2Key: Atan2,

		AtanhKey: Atanh,

		CbrtKey: Cbrt,

		CeilKey: Ceil,

		CopysignKey: Copysign,

		CosKey: Cos,

		CoshKey: Cosh,

		DimKey: Dim,

		ErfKey: Erf,

		ErfcKey: Erfc,

		ErfcinvKey: Erfcinv,

		ErfinvKey: Erfinv,

		ExpKey: Exp,

		Exp2Key: Exp2,

		Expm1Key: Expm1,

		Float32bitsKey: Float32bits,

		Float32frombitsKey: Float32frombits,

		Float64bitsKey: Float64bits,

		Float64frombitsKey: Float64frombits,

		FloorKey: Floor,

		FrexpKey: Frexp,

		GammaKey: Gamma,

		HypotKey: Hypot,

		IlogbKey: Ilogb,

		InfKey: Inf,

		IsInfKey: IsInf,

		IsNaNKey: IsNaN,

		J0Key: J0,

		J1Key: J1,

		JnKey: Jn,

		LdexpKey: Ldexp,

		LgammaKey: Lgamma,

		LogKey: Log,

		Log10Key: Log10,

		Log1pKey: Log1p,

		Log2Key: Log2,

		LogbKey: Logb,

		MaxKey: Max,

		MinKey: Min,

		ModKey: Mod,

		ModfKey: Modf,

		NaNKey: NaN,

		NextafterKey: Nextafter,

		Nextafter32Key: Nextafter32,

		PowKey: Pow,

		Pow10Key: Pow10,

		RemainderKey: Remainder,

		RoundKey: Round,

		RoundToEvenKey: RoundToEven,

		SignbitKey: Signbit,

		SinKey: Sin,

		SincosKey: Sincos,

		SinhKey: Sinh,

		SqrtKey: Sqrt,

		TanKey: Tan,

		TanhKey: Tanh,

		TruncKey: Trunc,

		Y0Key: Y0,

		Y1Key: Y1,

		YnKey: Yn,
	}
}

// Abs returns the absolute value of x.
//
// Special cases are:
// 	Abs(±Inf) = +Inf
// 	Abs(NaN) = NaN
var Abs = math.Abs

// Acos returns the arccosine, in radians, of x.
//
// Special case is:
// 	Acos(x) = NaN if x &lt; -1 or x &gt; 1
var Acos = math.Acos

// Acosh returns the inverse hyperbolic cosine of x.
//
// Special cases are:
// 	Acosh(+Inf) = +Inf
// 	Acosh(x) = NaN if x &lt; 1
// 	Acosh(NaN) = NaN
var Acosh = math.Acosh

// Asin returns the arcsine, in radians, of x.
//
// Special cases are:
// 	Asin(±0) = ±0
// 	Asin(x) = NaN if x &lt; -1 or x &gt; 1
var Asin = math.Asin

// Asinh returns the inverse hyperbolic sine of x.
//
// Special cases are:
// 	Asinh(±0) = ±0
// 	Asinh(±Inf) = ±Inf
// 	Asinh(NaN) = NaN
var Asinh = math.Asinh

// Atan returns the arctangent, in radians, of x.
//
// Special cases are:
//      Atan(±0) = ±0
//      Atan(±Inf) = ±Pi/2
var Atan = math.Atan

// Atan2 returns the arc tangent of y/x, using
// the signs of the two to determine the quadrant
// of the return value.
//
// Special cases are (in order):
// 	Atan2(y, NaN) = NaN
// 	Atan2(NaN, x) = NaN
// 	Atan2(+0, x&gt;=0) = +0
// 	Atan2(-0, x&gt;=0) = -0
// 	Atan2(+0, x&lt;=-0) = +Pi
// 	Atan2(-0, x&lt;=-0) = -Pi
// 	Atan2(y&gt;0, 0) = +Pi/2
// 	Atan2(y&lt;0, 0) = -Pi/2
// 	Atan2(+Inf, +Inf) = +Pi/4
// 	Atan2(-Inf, +Inf) = -Pi/4
// 	Atan2(+Inf, -Inf) = 3Pi/4
// 	Atan2(-Inf, -Inf) = -3Pi/4
// 	Atan2(y, +Inf) = 0
// 	Atan2(y&gt;0, -Inf) = +Pi
// 	Atan2(y&lt;0, -Inf) = -Pi
// 	Atan2(+Inf, x) = +Pi/2
// 	Atan2(-Inf, x) = -Pi/2
var Atan2 = math.Atan2

// Atanh returns the inverse hyperbolic tangent of x.
//
// Special cases are:
// 	Atanh(1) = +Inf
// 	Atanh(±0) = ±0
// 	Atanh(-1) = -Inf
// 	Atanh(x) = NaN if x &lt; -1 or x &gt; 1
// 	Atanh(NaN) = NaN
var Atanh = math.Atanh

// Cbrt returns the cube root of x.
//
// Special cases are:
// 	Cbrt(±0) = ±0
// 	Cbrt(±Inf) = ±Inf
// 	Cbrt(NaN) = NaN
var Cbrt = math.Cbrt

// Ceil returns the least integer value greater than or equal to x.
//
// Special cases are:
// 	Ceil(±0) = ±0
// 	Ceil(±Inf) = ±Inf
// 	Ceil(NaN) = NaN
var Ceil = math.Ceil

// Copysign returns a value with the magnitude
// of x and the sign of y.
var Copysign = math.Copysign

// Cos returns the cosine of the radian argument x.
//
// Special cases are:
// 	Cos(±Inf) = NaN
// 	Cos(NaN) = NaN
var Cos = math.Cos

// Cosh returns the hyperbolic cosine of x.
//
// Special cases are:
// 	Cosh(±0) = 1
// 	Cosh(±Inf) = +Inf
// 	Cosh(NaN) = NaN
var Cosh = math.Cosh

// Dim returns the maximum of x-y or 0.
//
// Special cases are:
// 	Dim(+Inf, +Inf) = NaN
// 	Dim(-Inf, -Inf) = NaN
// 	Dim(x, NaN) = Dim(NaN, x) = NaN
var Dim = math.Dim

// Erf returns the error function of x.
//
// Special cases are:
// 	Erf(+Inf) = 1
// 	Erf(-Inf) = -1
// 	Erf(NaN) = NaN
var Erf = math.Erf

// Erfc returns the complementary error function of x.
//
// Special cases are:
// 	Erfc(+Inf) = 0
// 	Erfc(-Inf) = 2
// 	Erfc(NaN) = NaN
var Erfc = math.Erfc

// Erfcinv returns the inverse of Erfc(x).
//
// Special cases are:
// 	Erfcinv(0) = +Inf
// 	Erfcinv(2) = -Inf
// 	Erfcinv(x) = NaN if x &lt; 0 or x &gt; 2
// 	Erfcinv(NaN) = NaN
var Erfcinv = math.Erfcinv

// Erfinv returns the inverse error function of x.
//
// Special cases are:
// 	Erfinv(1) = +Inf
// 	Erfinv(-1) = -Inf
// 	Erfinv(x) = NaN if x &lt; -1 or x &gt; 1
// 	Erfinv(NaN) = NaN
var Erfinv = math.Erfinv

// Exp returns e**x, the base-e exponential of x.
//
// Special cases are:
// 	Exp(+Inf) = +Inf
// 	Exp(NaN) = NaN
// Very large values overflow to 0 or +Inf.
// Very small values underflow to 1.
var Exp = math.Exp

// Exp2 returns 2**x, the base-2 exponential of x.
//
// Special cases are the same as Exp.
var Exp2 = math.Exp2

// Expm1 returns e**x - 1, the base-e exponential of x minus 1.
// It is more accurate than Exp(x) - 1 when x is near zero.
//
// Special cases are:
// 	Expm1(+Inf) = +Inf
// 	Expm1(-Inf) = -1
// 	Expm1(NaN) = NaN
// Very large values overflow to -1 or +Inf.
var Expm1 = math.Expm1

// Float32bits returns the IEEE 754 binary representation of f.
var Float32bits = math.Float32bits

// Float32frombits returns the floating point number corresponding
// to the IEEE 754 binary representation b.
var Float32frombits = math.Float32frombits

// Float64bits returns the IEEE 754 binary representation of f.
var Float64bits = math.Float64bits

// Float64frombits returns the floating point number corresponding
// the IEEE 754 binary representation b.
var Float64frombits = math.Float64frombits

// Floor returns the greatest integer value less than or equal to x.
//
// Special cases are:
// 	Floor(±0) = ±0
// 	Floor(±Inf) = ±Inf
// 	Floor(NaN) = NaN
var Floor = math.Floor

// Frexp breaks f into a normalized fraction
// and an integral power of two.
// It returns frac and exp satisfying f == frac × 2**exp,
// with the absolute value of frac in the interval [½, 1).
//
// Special cases are:
// 	Frexp(±0) = ±0, 0
// 	Frexp(±Inf) = ±Inf, 0
// 	Frexp(NaN) = NaN, 0
var Frexp = math.Frexp

// Gamma returns the Gamma function of x.
//
// Special cases are:
// 	Gamma(+Inf) = +Inf
// 	Gamma(+0) = +Inf
// 	Gamma(-0) = -Inf
// 	Gamma(x) = NaN for integer x &lt; 0
// 	Gamma(-Inf) = NaN
// 	Gamma(NaN) = NaN
var Gamma = math.Gamma

// Hypot returns Sqrt(p*p + q*q), taking care to avoid
// unnecessary overflow and underflow.
//
// Special cases are:
// 	Hypot(±Inf, q) = +Inf
// 	Hypot(p, ±Inf) = +Inf
// 	Hypot(NaN, q) = NaN
// 	Hypot(p, NaN) = NaN
var Hypot = math.Hypot

// Ilogb returns the binary exponent of x as an integer.
//
// Special cases are:
// 	Ilogb(±Inf) = MaxInt32
// 	Ilogb(0) = MinInt32
// 	Ilogb(NaN) = MaxInt32
var Ilogb = math.Ilogb

// Inf returns positive infinity if sign &gt;= 0, negative infinity if sign &lt; 0.
var Inf = math.Inf

// IsInf reports whether f is an infinity, according to sign.
// If sign &gt; 0, IsInf reports whether f is positive infinity.
// If sign &lt; 0, IsInf reports whether f is negative infinity.
// If sign == 0, IsInf reports whether f is either infinity.
var IsInf = math.IsInf

// IsNaN reports whether f is an IEEE 754 ``not-a-number&#39;&#39; value.
var IsNaN = math.IsNaN

// J0 returns the order-zero Bessel function of the first kind.
//
// Special cases are:
// 	J0(±Inf) = 0
// 	J0(0) = 1
// 	J0(NaN) = NaN
var J0 = math.J0

// J1 returns the order-one Bessel function of the first kind.
//
// Special cases are:
// 	J1(±Inf) = 0
// 	J1(NaN) = NaN
var J1 = math.J1

// Jn returns the order-n Bessel function of the first kind.
//
// Special cases are:
// 	Jn(n, ±Inf) = 0
// 	Jn(n, NaN) = NaN
var Jn = math.Jn

// Ldexp is the inverse of Frexp.
// It returns frac × 2**exp.
//
// Special cases are:
// 	Ldexp(±0, exp) = ±0
// 	Ldexp(±Inf, exp) = ±Inf
// 	Ldexp(NaN, exp) = NaN
var Ldexp = math.Ldexp

// Lgamma returns the natural logarithm and sign (-1 or +1) of Gamma(x).
//
// Special cases are:
// 	Lgamma(+Inf) = +Inf
// 	Lgamma(0) = +Inf
// 	Lgamma(-integer) = +Inf
// 	Lgamma(-Inf) = -Inf
// 	Lgamma(NaN) = NaN
var Lgamma = math.Lgamma

// Log returns the natural logarithm of x.
//
// Special cases are:
// 	Log(+Inf) = +Inf
// 	Log(0) = -Inf
// 	Log(x &lt; 0) = NaN
// 	Log(NaN) = NaN
var Log = math.Log

// Log10 returns the decimal logarithm of x.
// The special cases are the same as for Log.
var Log10 = math.Log10

// Log1p returns the natural logarithm of 1 plus its argument x.
// It is more accurate than Log(1 + x) when x is near zero.
//
// Special cases are:
// 	Log1p(+Inf) = +Inf
// 	Log1p(±0) = ±0
// 	Log1p(-1) = -Inf
// 	Log1p(x &lt; -1) = NaN
// 	Log1p(NaN) = NaN
var Log1p = math.Log1p

// Log2 returns the binary logarithm of x.
// The special cases are the same as for Log.
var Log2 = math.Log2

// Logb returns the binary exponent of x.
//
// Special cases are:
// 	Logb(±Inf) = +Inf
// 	Logb(0) = -Inf
// 	Logb(NaN) = NaN
var Logb = math.Logb

// Max returns the larger of x or y.
//
// Special cases are:
// 	Max(x, +Inf) = Max(+Inf, x) = +Inf
// 	Max(x, NaN) = Max(NaN, x) = NaN
// 	Max(+0, ±0) = Max(±0, +0) = +0
// 	Max(-0, -0) = -0
var Max = math.Max

// Min returns the smaller of x or y.
//
// Special cases are:
// 	Min(x, -Inf) = Min(-Inf, x) = -Inf
// 	Min(x, NaN) = Min(NaN, x) = NaN
// 	Min(-0, ±0) = Min(±0, -0) = -0
var Min = math.Min

// Mod returns the floating-point remainder of x/y.
// The magnitude of the result is less than y and its
// sign agrees with that of x.
//
// Special cases are:
// 	Mod(±Inf, y) = NaN
// 	Mod(NaN, y) = NaN
// 	Mod(x, 0) = NaN
// 	Mod(x, ±Inf) = x
// 	Mod(x, NaN) = NaN
var Mod = math.Mod

// Modf returns integer and fractional floating-point numbers
// that sum to f. Both values have the same sign as f.
//
// Special cases are:
// 	Modf(±Inf) = ±Inf, NaN
// 	Modf(NaN) = NaN, NaN
var Modf = math.Modf

// NaN returns an IEEE 754 ``not-a-number&#39;&#39; value.
var NaN = math.NaN

// Nextafter returns the next representable float64 value after x towards y.
//
// Special cases are:
// 	Nextafter(x, x)   = x
// 	Nextafter(NaN, y) = NaN
// 	Nextafter(x, NaN) = NaN
var Nextafter = math.Nextafter

// Nextafter32 returns the next representable float32 value after x towards y.
//
// Special cases are:
// 	Nextafter32(x, x)   = x
// 	Nextafter32(NaN, y) = NaN
// 	Nextafter32(x, NaN) = NaN
var Nextafter32 = math.Nextafter32

// Pow returns x**y, the base-x exponential of y.
//
// Special cases are (in order):
// 	Pow(x, ±0) = 1 for any x
// 	Pow(1, y) = 1 for any y
// 	Pow(x, 1) = x for any x
// 	Pow(NaN, y) = NaN
// 	Pow(x, NaN) = NaN
// 	Pow(±0, y) = ±Inf for y an odd integer &lt; 0
// 	Pow(±0, -Inf) = +Inf
// 	Pow(±0, +Inf) = +0
// 	Pow(±0, y) = +Inf for finite y &lt; 0 and not an odd integer
// 	Pow(±0, y) = ±0 for y an odd integer &gt; 0
// 	Pow(±0, y) = +0 for finite y &gt; 0 and not an odd integer
// 	Pow(-1, ±Inf) = 1
// 	Pow(x, +Inf) = +Inf for |x| &gt; 1
// 	Pow(x, -Inf) = +0 for |x| &gt; 1
// 	Pow(x, +Inf) = +0 for |x| &lt; 1
// 	Pow(x, -Inf) = +Inf for |x| &lt; 1
// 	Pow(+Inf, y) = +Inf for y &gt; 0
// 	Pow(+Inf, y) = +0 for y &lt; 0
// 	Pow(-Inf, y) = Pow(-0, -y)
// 	Pow(x, y) = NaN for finite x &lt; 0 and finite non-integer y
var Pow = math.Pow

// Pow10 returns 10**n, the base-10 exponential of n.
//
// Special cases are:
// 	Pow10(n) =    0 for n &lt; -323
// 	Pow10(n) = +Inf for n &gt; 308
var Pow10 = math.Pow10

// Remainder returns the IEEE 754 floating-point remainder of x/y.
//
// Special cases are:
// 	Remainder(±Inf, y) = NaN
// 	Remainder(NaN, y) = NaN
// 	Remainder(x, 0) = NaN
// 	Remainder(x, ±Inf) = x
// 	Remainder(x, NaN) = NaN
var Remainder = math.Remainder

// Round returns the nearest integer, rounding half away from zero.
//
// Special cases are:
// 	Round(±0) = ±0
// 	Round(±Inf) = ±Inf
// 	Round(NaN) = NaN
var Round = math.Round

// RoundToEven returns the nearest integer, rounding ties to even.
//
// Special cases are:
// 	RoundToEven(±0) = ±0
// 	RoundToEven(±Inf) = ±Inf
// 	RoundToEven(NaN) = NaN
var RoundToEven = math.RoundToEven

// Signbit returns true if x is negative or negative zero.
var Signbit = math.Signbit

// Sin returns the sine of the radian argument x.
//
// Special cases are:
// 	Sin(±0) = ±0
// 	Sin(±Inf) = NaN
// 	Sin(NaN) = NaN
var Sin = math.Sin

// Sincos returns Sin(x), Cos(x).
//
// Special cases are:
// 	Sincos(±0) = ±0, 1
// 	Sincos(±Inf) = NaN, NaN
// 	Sincos(NaN) = NaN, NaN
var Sincos = math.Sincos

// Sinh returns the hyperbolic sine of x.
//
// Special cases are:
// 	Sinh(±0) = ±0
// 	Sinh(±Inf) = ±Inf
// 	Sinh(NaN) = NaN
var Sinh = math.Sinh

// Sqrt returns the square root of x.
//
// Special cases are:
// 	Sqrt(+Inf) = +Inf
// 	Sqrt(±0) = ±0
// 	Sqrt(x &lt; 0) = NaN
// 	Sqrt(NaN) = NaN
var Sqrt = math.Sqrt

// Tan returns the tangent of the radian argument x.
//
// Special cases are:
// 	Tan(±0) = ±0
// 	Tan(±Inf) = NaN
// 	Tan(NaN) = NaN
var Tan = math.Tan

// Tanh returns the hyperbolic tangent of x.
//
// Special cases are:
// 	Tanh(±0) = ±0
// 	Tanh(±Inf) = ±1
// 	Tanh(NaN) = NaN
var Tanh = math.Tanh

// Trunc returns the integer value of x.
//
// Special cases are:
// 	Trunc(±0) = ±0
// 	Trunc(±Inf) = ±Inf
// 	Trunc(NaN) = NaN
var Trunc = math.Trunc

// Y0 returns the order-zero Bessel function of the second kind.
//
// Special cases are:
// 	Y0(+Inf) = 0
// 	Y0(0) = -Inf
// 	Y0(x &lt; 0) = NaN
// 	Y0(NaN) = NaN
var Y0 = math.Y0

// Y1 returns the order-one Bessel function of the second kind.
//
// Special cases are:
// 	Y1(+Inf) = 0
// 	Y1(0) = -Inf
// 	Y1(x &lt; 0) = NaN
// 	Y1(NaN) = NaN
var Y1 = math.Y1

// Yn returns the order-n Bessel function of the second kind.
//
// Special cases are:
// 	Yn(n, +Inf) = 0
// 	Yn(n ≥ 0, 0) = -Inf
// 	Yn(n &lt; 0, 0) = +Inf if n is odd, -Inf if n is even
// 	Yn(n, x &lt; 0) = NaN
// 	Yn(n, NaN) = NaN
var Yn = math.Yn
