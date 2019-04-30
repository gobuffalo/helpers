package randh

import (
	"math/rand"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ExpFloat64Key = "ExpFloat64"

	Float32Key = "Float32"

	Float64Key = "Float64"

	IntKey = "Int"

	Int31Key = "Int31"

	Int31nKey = "Int31n"

	Int63Key = "Int63"

	Int63nKey = "Int63n"

	IntnKey = "Intn"

	NewKey = "New"

	NewSourceKey = "NewSource"

	NewZipfKey = "NewZipf"

	NormFloat64Key = "NormFloat64"

	PermKey = "Perm"

	ReadKey = "Read"

	SeedKey = "Seed"

	ShuffleKey = "Shuffle"

	Uint32Key = "Uint32"

	Uint64Key = "Uint64"
)

func New() hctx.Map {
	return hctx.Map{

		ExpFloat64Key: ExpFloat64,

		Float32Key: Float32,

		Float64Key: Float64,

		IntKey: Int,

		Int31Key: Int31,

		Int31nKey: Int31n,

		Int63Key: Int63,

		Int63nKey: Int63n,

		IntnKey: Intn,

		NewKey: New,

		NewSourceKey: NewSource,

		NewZipfKey: NewZipf,

		NormFloat64Key: NormFloat64,

		PermKey: Perm,

		ReadKey: Read,

		SeedKey: Seed,

		ShuffleKey: Shuffle,

		Uint32Key: Uint32,

		Uint64Key: Uint64,
	}
}

// ExpFloat64 returns an exponentially distributed float64 in the range
// (0, +math.MaxFloat64] with an exponential distribution whose rate parameter
// (lambda) is 1 and whose mean is 1/lambda (1) from the default Source.
// To produce a distribution with a different rate parameter,
// callers can adjust the output using:
//
//  sample = ExpFloat64() / desiredRateParameter
var ExpFloat64 = rand.ExpFloat64

// Float32 returns, as a float32, a pseudo-random number in [0.0,1.0)
// from the default Source.
var Float32 = rand.Float32

// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0)
// from the default Source.
var Float64 = rand.Float64

// Int returns a non-negative pseudo-random int from the default Source.
var Int = rand.Int

// Int31 returns a non-negative pseudo-random 31-bit integer as an int32
// from the default Source.
var Int31 = rand.Int31

// Int31n returns, as an int32, a non-negative pseudo-random number in [0,n)
// from the default Source.
// It panics if n &lt;= 0.
var Int31n = rand.Int31n

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64
// from the default Source.
var Int63 = rand.Int63

// Int63n returns, as an int64, a non-negative pseudo-random number in [0,n)
// from the default Source.
// It panics if n &lt;= 0.
var Int63n = rand.Int63n

// Intn returns, as an int, a non-negative pseudo-random number in [0,n)
// from the default Source.
// It panics if n &lt;= 0.
var Intn = rand.Intn

// New returns a new Rand that uses random values from src
// to generate other random values.
var New = rand.New

// NewSource returns a new pseudo-random Source seeded with the given value.
// Unlike the default Source used by top-level functions, this source is not
// safe for concurrent use by multiple goroutines.
var NewSource = rand.NewSource

// NewZipf returns a Zipf variate generator.
// The generator generates values k ∈ [0, imax]
// such that P(k) is proportional to (v + k) ** (-s).
// Requirements: s &gt; 1 and v &gt;= 1.
var NewZipf = rand.NewZipf

// NormFloat64 returns a normally distributed float64 in the range
// [-math.MaxFloat64, +math.MaxFloat64] with
// standard normal distribution (mean = 0, stddev = 1)
// from the default Source.
// To produce a different normal distribution, callers can
// adjust the output using:
//
//  sample = NormFloat64() * desiredStdDev + desiredMean
var NormFloat64 = rand.NormFloat64

// Perm returns, as a slice of n ints, a pseudo-random permutation of the integers [0,n)
// from the default Source.
var Perm = rand.Perm

// Read generates len(p) random bytes from the default Source and
// writes them into p. It always returns len(p) and a nil error.
// Read, unlike the Rand.Read method, is safe for concurrent use.
var Read = rand.Read

// Seed uses the provided seed value to initialize the default Source to a
// deterministic state. If Seed is not called, the generator behaves as
// if seeded by Seed(1). Seed values that have the same remainder when
// divided by 2^31-1 generate the same pseudo-random sequence.
// Seed, unlike the Rand.Seed method, is safe for concurrent use.
var Seed = rand.Seed

// Shuffle pseudo-randomizes the order of elements using the default Source.
// n is the number of elements. Shuffle panics if n &lt; 0.
// swap swaps the elements with indexes i and j.
var Shuffle = rand.Shuffle

// Uint32 returns a pseudo-random 32-bit value as a uint32
// from the default Source.
var Uint32 = rand.Uint32

// Uint64 returns a pseudo-random 64-bit value as a uint64
// from the default Source.
var Uint64 = rand.Uint64
