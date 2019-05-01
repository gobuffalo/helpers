package cmplxh

import (
	"math/cmplx"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AbsKey = "Abs"

	AcosKey = "Acos"

	AcoshKey = "Acosh"

	AsinKey = "Asin"

	AsinhKey = "Asinh"

	AtanKey = "Atan"

	AtanhKey = "Atanh"

	ConjKey = "Conj"

	CosKey = "Cos"

	CoshKey = "Cosh"

	CotKey = "Cot"

	ExpKey = "Exp"

	InfKey = "Inf"

	IsInfKey = "IsInf"

	IsNaNKey = "IsNaN"

	LogKey = "Log"

	Log10Key = "Log10"

	NaNKey = "NaN"

	PhaseKey = "Phase"

	PolarKey = "Polar"

	PowKey = "Pow"

	RectKey = "Rect"

	SinKey = "Sin"

	SinhKey = "Sinh"

	SqrtKey = "Sqrt"

	TanKey = "Tan"

	TanhKey = "Tanh"
)

func New() hctx.Map {
	return hctx.Map{

		AbsKey: Abs,

		AcosKey: Acos,

		AcoshKey: Acosh,

		AsinKey: Asin,

		AsinhKey: Asinh,

		AtanKey: Atan,

		AtanhKey: Atanh,

		ConjKey: Conj,

		CosKey: Cos,

		CoshKey: Cosh,

		CotKey: Cot,

		ExpKey: Exp,

		InfKey: Inf,

		IsInfKey: IsInf,

		IsNaNKey: IsNaN,

		LogKey: Log,

		Log10Key: Log10,

		NaNKey: NaN,

		PhaseKey: Phase,

		PolarKey: Polar,

		PowKey: Pow,

		RectKey: Rect,

		SinKey: Sin,

		SinhKey: Sinh,

		SqrtKey: Sqrt,

		TanKey: Tan,

		TanhKey: Tanh,
	}
}

// Abs returns the absolute value (also called the modulus) of x.
var Abs = cmplx.Abs

// Acos returns the inverse cosine of x.
var Acos = cmplx.Acos

// Acosh returns the inverse hyperbolic cosine of x.
var Acosh = cmplx.Acosh

// Asin returns the inverse sine of x.
var Asin = cmplx.Asin

// Asinh returns the inverse hyperbolic sine of x.
var Asinh = cmplx.Asinh

// Atan returns the inverse tangent of x.
var Atan = cmplx.Atan

// Atanh returns the inverse hyperbolic tangent of x.
var Atanh = cmplx.Atanh

// Conj returns the complex conjugate of x.
var Conj = cmplx.Conj

// Cos returns the cosine of x.
var Cos = cmplx.Cos

// Cosh returns the hyperbolic cosine of x.
var Cosh = cmplx.Cosh

// Cot returns the cotangent of x.
var Cot = cmplx.Cot

// Exp returns e**x, the base-e exponential of x.
var Exp = cmplx.Exp

// Inf returns a complex infinity, complex(+Inf, +Inf).
var Inf = cmplx.Inf

// IsInf returns true if either real(x) or imag(x) is an infinity.
var IsInf = cmplx.IsInf

// IsNaN returns true if either real(x) or imag(x) is NaN
// and neither is an infinity.
var IsNaN = cmplx.IsNaN

// Log returns the natural logarithm of x.
var Log = cmplx.Log

// Log10 returns the decimal logarithm of x.
var Log10 = cmplx.Log10

// NaN returns a complex ``not-a-number&#39;&#39; value.
var NaN = cmplx.NaN

// Phase returns the phase (also called the argument) of x.
// The returned value is in the range [-Pi, Pi].
var Phase = cmplx.Phase

// Polar returns the absolute value r and phase θ of x,
// such that x = r * e**θi.
// The phase is in the range [-Pi, Pi].
var Polar = cmplx.Polar

// Pow returns x**y, the base-x exponential of y.
// For generalized compatibility with math.Pow:
// 	Pow(0, ±0) returns 1+0i
// 	Pow(0, c) for real(c)&lt;0 returns Inf+0i if imag(c) is zero, otherwise Inf+Inf i.
var Pow = cmplx.Pow

// Rect returns the complex number x with polar coordinates r, θ.
var Rect = cmplx.Rect

// Sin returns the sine of x.
var Sin = cmplx.Sin

// Sinh returns the hyperbolic sine of x.
var Sinh = cmplx.Sinh

// Sqrt returns the square root of x.
// The result r is chosen so that real(r) ≥ 0 and imag(r) has the same sign as imag(x).
var Sqrt = cmplx.Sqrt

// Tan returns the tangent of x.
var Tan = cmplx.Tan

// Tanh returns the hyperbolic tangent of x.
var Tanh = cmplx.Tanh
