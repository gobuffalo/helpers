package bigh

import (
	"math/big"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	JacobiKey = "Jacobi"

	NewFloatKey = "NewFloat"

	NewIntKey = "NewInt"

	NewRatKey = "NewRat"

	ParseFloatKey = "ParseFloat"
)

func New() hctx.Map {
	return hctx.Map{

		JacobiKey: Jacobi,

		NewFloatKey: NewFloat,

		NewIntKey: NewInt,

		NewRatKey: NewRat,

		ParseFloatKey: ParseFloat,
	}
}

// Jacobi returns the Jacobi symbol (x/y), either +1, -1, or 0.
// The y argument must be an odd integer.
var Jacobi = big.Jacobi

// NewFloat allocates and returns a new Float set to x,
// with precision 53 and rounding mode ToNearestEven.
// NewFloat panics with ErrNaN if x is a NaN.
var NewFloat = big.NewFloat

// NewInt allocates and returns a new Int set to x.
var NewInt = big.NewInt

// NewRat creates a new Rat with numerator a and denominator b.
var NewRat = big.NewRat

// ParseFloat is like f.Parse(s, base) with f set to the given precision
// and rounding mode.
var ParseFloat = big.ParseFloat
