package bitsh

import (
	"math/bits"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	LeadingZerosKey = "LeadingZeros"

	LeadingZeros16Key = "LeadingZeros16"

	LeadingZeros32Key = "LeadingZeros32"

	LeadingZeros64Key = "LeadingZeros64"

	LeadingZeros8Key = "LeadingZeros8"

	LenKey = "Len"

	Len16Key = "Len16"

	Len32Key = "Len32"

	Len64Key = "Len64"

	Len8Key = "Len8"

	OnesCountKey = "OnesCount"

	OnesCount16Key = "OnesCount16"

	OnesCount32Key = "OnesCount32"

	OnesCount64Key = "OnesCount64"

	OnesCount8Key = "OnesCount8"

	ReverseKey = "Reverse"

	Reverse16Key = "Reverse16"

	Reverse32Key = "Reverse32"

	Reverse64Key = "Reverse64"

	Reverse8Key = "Reverse8"

	ReverseBytesKey = "ReverseBytes"

	ReverseBytes16Key = "ReverseBytes16"

	ReverseBytes32Key = "ReverseBytes32"

	ReverseBytes64Key = "ReverseBytes64"

	RotateLeftKey = "RotateLeft"

	RotateLeft16Key = "RotateLeft16"

	RotateLeft32Key = "RotateLeft32"

	RotateLeft64Key = "RotateLeft64"

	RotateLeft8Key = "RotateLeft8"

	TrailingZerosKey = "TrailingZeros"

	TrailingZeros16Key = "TrailingZeros16"

	TrailingZeros32Key = "TrailingZeros32"

	TrailingZeros64Key = "TrailingZeros64"

	TrailingZeros8Key = "TrailingZeros8"
)

func New() hctx.Map {
	return hctx.Map{

		LeadingZerosKey: LeadingZeros,

		LeadingZeros16Key: LeadingZeros16,

		LeadingZeros32Key: LeadingZeros32,

		LeadingZeros64Key: LeadingZeros64,

		LeadingZeros8Key: LeadingZeros8,

		LenKey: Len,

		Len16Key: Len16,

		Len32Key: Len32,

		Len64Key: Len64,

		Len8Key: Len8,

		OnesCountKey: OnesCount,

		OnesCount16Key: OnesCount16,

		OnesCount32Key: OnesCount32,

		OnesCount64Key: OnesCount64,

		OnesCount8Key: OnesCount8,

		ReverseKey: Reverse,

		Reverse16Key: Reverse16,

		Reverse32Key: Reverse32,

		Reverse64Key: Reverse64,

		Reverse8Key: Reverse8,

		ReverseBytesKey: ReverseBytes,

		ReverseBytes16Key: ReverseBytes16,

		ReverseBytes32Key: ReverseBytes32,

		ReverseBytes64Key: ReverseBytes64,

		RotateLeftKey: RotateLeft,

		RotateLeft16Key: RotateLeft16,

		RotateLeft32Key: RotateLeft32,

		RotateLeft64Key: RotateLeft64,

		RotateLeft8Key: RotateLeft8,

		TrailingZerosKey: TrailingZeros,

		TrailingZeros16Key: TrailingZeros16,

		TrailingZeros32Key: TrailingZeros32,

		TrailingZeros64Key: TrailingZeros64,

		TrailingZeros8Key: TrailingZeros8,
	}
}

// LeadingZeros returns the number of leading zero bits in x; the result is UintSize for x == 0.
var LeadingZeros = bits.LeadingZeros

// LeadingZeros16 returns the number of leading zero bits in x; the result is 16 for x == 0.
var LeadingZeros16 = bits.LeadingZeros16

// LeadingZeros32 returns the number of leading zero bits in x; the result is 32 for x == 0.
var LeadingZeros32 = bits.LeadingZeros32

// LeadingZeros64 returns the number of leading zero bits in x; the result is 64 for x == 0.
var LeadingZeros64 = bits.LeadingZeros64

// LeadingZeros8 returns the number of leading zero bits in x; the result is 8 for x == 0.
var LeadingZeros8 = bits.LeadingZeros8

// Len returns the minimum number of bits required to represent x; the result is 0 for x == 0.
var Len = bits.Len

// Len16 returns the minimum number of bits required to represent x; the result is 0 for x == 0.
var Len16 = bits.Len16

// Len32 returns the minimum number of bits required to represent x; the result is 0 for x == 0.
var Len32 = bits.Len32

// Len64 returns the minimum number of bits required to represent x; the result is 0 for x == 0.
var Len64 = bits.Len64

// Len8 returns the minimum number of bits required to represent x; the result is 0 for x == 0.
var Len8 = bits.Len8

// OnesCount returns the number of one bits (&#34;population count&#34;) in x.
var OnesCount = bits.OnesCount

// OnesCount16 returns the number of one bits (&#34;population count&#34;) in x.
var OnesCount16 = bits.OnesCount16

// OnesCount32 returns the number of one bits (&#34;population count&#34;) in x.
var OnesCount32 = bits.OnesCount32

// OnesCount64 returns the number of one bits (&#34;population count&#34;) in x.
var OnesCount64 = bits.OnesCount64

// OnesCount8 returns the number of one bits (&#34;population count&#34;) in x.
var OnesCount8 = bits.OnesCount8

// Reverse returns the value of x with its bits in reversed order.
var Reverse = bits.Reverse

// Reverse16 returns the value of x with its bits in reversed order.
var Reverse16 = bits.Reverse16

// Reverse32 returns the value of x with its bits in reversed order.
var Reverse32 = bits.Reverse32

// Reverse64 returns the value of x with its bits in reversed order.
var Reverse64 = bits.Reverse64

// Reverse8 returns the value of x with its bits in reversed order.
var Reverse8 = bits.Reverse8

// ReverseBytes returns the value of x with its bytes in reversed order.
var ReverseBytes = bits.ReverseBytes

// ReverseBytes16 returns the value of x with its bytes in reversed order.
var ReverseBytes16 = bits.ReverseBytes16

// ReverseBytes32 returns the value of x with its bytes in reversed order.
var ReverseBytes32 = bits.ReverseBytes32

// ReverseBytes64 returns the value of x with its bytes in reversed order.
var ReverseBytes64 = bits.ReverseBytes64

// RotateLeft returns the value of x rotated left by (k mod UintSize) bits.
// To rotate x right by k bits, call RotateLeft(x, -k).
var RotateLeft = bits.RotateLeft

// RotateLeft16 returns the value of x rotated left by (k mod 16) bits.
// To rotate x right by k bits, call RotateLeft16(x, -k).
var RotateLeft16 = bits.RotateLeft16

// RotateLeft32 returns the value of x rotated left by (k mod 32) bits.
// To rotate x right by k bits, call RotateLeft32(x, -k).
var RotateLeft32 = bits.RotateLeft32

// RotateLeft64 returns the value of x rotated left by (k mod 64) bits.
// To rotate x right by k bits, call RotateLeft64(x, -k).
var RotateLeft64 = bits.RotateLeft64

// RotateLeft8 returns the value of x rotated left by (k mod 8) bits.
// To rotate x right by k bits, call RotateLeft8(x, -k).
var RotateLeft8 = bits.RotateLeft8

// TrailingZeros returns the number of trailing zero bits in x; the result is UintSize for x == 0.
var TrailingZeros = bits.TrailingZeros

// TrailingZeros16 returns the number of trailing zero bits in x; the result is 16 for x == 0.
var TrailingZeros16 = bits.TrailingZeros16

// TrailingZeros32 returns the number of trailing zero bits in x; the result is 32 for x == 0.
var TrailingZeros32 = bits.TrailingZeros32

// TrailingZeros64 returns the number of trailing zero bits in x; the result is 64 for x == 0.
var TrailingZeros64 = bits.TrailingZeros64

// TrailingZeros8 returns the number of trailing zero bits in x; the result is 8 for x == 0.
var TrailingZeros8 = bits.TrailingZeros8
