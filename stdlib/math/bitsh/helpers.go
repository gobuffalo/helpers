package bitsh

import (
	"math/bits"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AddKey = "Add"

	Add32Key = "Add32"

	Add64Key = "Add64"

	DivKey = "Div"

	Div32Key = "Div32"

	Div64Key = "Div64"

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

	MulKey = "Mul"

	Mul32Key = "Mul32"

	Mul64Key = "Mul64"

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

	SubKey = "Sub"

	Sub32Key = "Sub32"

	Sub64Key = "Sub64"

	TrailingZerosKey = "TrailingZeros"

	TrailingZeros16Key = "TrailingZeros16"

	TrailingZeros32Key = "TrailingZeros32"

	TrailingZeros64Key = "TrailingZeros64"

	TrailingZeros8Key = "TrailingZeros8"
)

func New() hctx.Map {
	return hctx.Map{

		AddKey: Add,

		Add32Key: Add32,

		Add64Key: Add64,

		DivKey: Div,

		Div32Key: Div32,

		Div64Key: Div64,

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

		MulKey: Mul,

		Mul32Key: Mul32,

		Mul64Key: Mul64,

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

		SubKey: Sub,

		Sub32Key: Sub32,

		Sub64Key: Sub64,

		TrailingZerosKey: TrailingZeros,

		TrailingZeros16Key: TrailingZeros16,

		TrailingZeros32Key: TrailingZeros32,

		TrailingZeros64Key: TrailingZeros64,

		TrailingZeros8Key: TrailingZeros8,
	}
}

// Add returns the sum with carry of x, y and carry: sum = x + y + carry.
// The carry input must be 0 or 1; otherwise the behavior is undefined.
// The carryOut output is guaranteed to be 0 or 1.
var Add = bits.Add

// Add32 returns the sum with carry of x, y and carry: sum = x + y + carry.
// The carry input must be 0 or 1; otherwise the behavior is undefined.
// The carryOut output is guaranteed to be 0 or 1.
var Add32 = bits.Add32

// Add64 returns the sum with carry of x, y and carry: sum = x + y + carry.
// The carry input must be 0 or 1; otherwise the behavior is undefined.
// The carryOut output is guaranteed to be 0 or 1.
var Add64 = bits.Add64

// Div returns the quotient and remainder of (hi, lo) divided by y:
// quo = (hi, lo)/y, rem = (hi, lo)%y with the dividend bits&#39; upper
// half in parameter hi and the lower half in parameter lo.
// Div panics for y == 0 (division by zero) or y &lt;= hi (quotient overflow).
var Div = bits.Div

// Div32 returns the quotient and remainder of (hi, lo) divided by y:
// quo = (hi, lo)/y, rem = (hi, lo)%y with the dividend bits&#39; upper
// half in parameter hi and the lower half in parameter lo.
// Div32 panics for y == 0 (division by zero) or y &lt;= hi (quotient overflow).
var Div32 = bits.Div32

// Div64 returns the quotient and remainder of (hi, lo) divided by y:
// quo = (hi, lo)/y, rem = (hi, lo)%y with the dividend bits&#39; upper
// half in parameter hi and the lower half in parameter lo.
// Div64 panics for y == 0 (division by zero) or y &lt;= hi (quotient overflow).
var Div64 = bits.Div64

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

// Mul returns the full-width product of x and y: (hi, lo) = x * y
// with the product bits&#39; upper half returned in hi and the lower
// half returned in lo.
var Mul = bits.Mul

// Mul32 returns the 64-bit product of x and y: (hi, lo) = x * y
// with the product bits&#39; upper half returned in hi and the lower
// half returned in lo.
var Mul32 = bits.Mul32

// Mul64 returns the 128-bit product of x and y: (hi, lo) = x * y
// with the product bits&#39; upper half returned in hi and the lower
// half returned in lo.
var Mul64 = bits.Mul64

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

// Sub returns the difference of x, y and borrow: diff = x - y - borrow.
// The borrow input must be 0 or 1; otherwise the behavior is undefined.
// The borrowOut output is guaranteed to be 0 or 1.
var Sub = bits.Sub

// Sub32 returns the difference of x, y and borrow, diff = x - y - borrow.
// The borrow input must be 0 or 1; otherwise the behavior is undefined.
// The borrowOut output is guaranteed to be 0 or 1.
var Sub32 = bits.Sub32

// Sub64 returns the difference of x, y and borrow: diff = x - y - borrow.
// The borrow input must be 0 or 1; otherwise the behavior is undefined.
// The borrowOut output is guaranteed to be 0 or 1.
var Sub64 = bits.Sub64

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
