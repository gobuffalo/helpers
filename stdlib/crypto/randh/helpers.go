package randh

import (
	"crypto/rand"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	IntKey = "Int"

	PrimeKey = "Prime"

	ReadKey = "Read"
)

func New() hctx.Map {
	return hctx.Map{

		IntKey: Int,

		PrimeKey: Prime,

		ReadKey: Read,
	}
}

// Int returns a uniform random value in [0, max). It panics if max &lt;= 0.
var Int = rand.Int

// Prime returns a number, p, of the given size, such that p is prime
// with high probability.
// Prime will return error for any error returned by rand.Read or if bits &lt; 2.
var Prime = rand.Prime

// Read is a helper function that calls Reader.Read using io.ReadFull.
// On return, n == len(b) if and only if err == nil.
var Read = rand.Read
