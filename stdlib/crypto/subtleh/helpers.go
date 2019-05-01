package subtleh

import (
	"crypto/subtle"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ConstantTimeByteEqKey = "ConstantTimeByteEq"

	ConstantTimeCompareKey = "ConstantTimeCompare"

	ConstantTimeCopyKey = "ConstantTimeCopy"

	ConstantTimeEqKey = "ConstantTimeEq"

	ConstantTimeLessOrEqKey = "ConstantTimeLessOrEq"

	ConstantTimeSelectKey = "ConstantTimeSelect"
)

func New() hctx.Map {
	return hctx.Map{

		ConstantTimeByteEqKey: ConstantTimeByteEq,

		ConstantTimeCompareKey: ConstantTimeCompare,

		ConstantTimeCopyKey: ConstantTimeCopy,

		ConstantTimeEqKey: ConstantTimeEq,

		ConstantTimeLessOrEqKey: ConstantTimeLessOrEq,

		ConstantTimeSelectKey: ConstantTimeSelect,
	}
}

// ConstantTimeByteEq returns 1 if x == y and 0 otherwise.
var ConstantTimeByteEq = subtle.ConstantTimeByteEq

// ConstantTimeCompare returns 1 if and only if the two slices, x
// and y, have equal contents. The time taken is a function of the length of
// the slices and is independent of the contents.
var ConstantTimeCompare = subtle.ConstantTimeCompare

// ConstantTimeCopy copies the contents of y into x (a slice of equal length)
// if v == 1. If v == 0, x is left unchanged. Its behavior is undefined if v
// takes any other value.
var ConstantTimeCopy = subtle.ConstantTimeCopy

// ConstantTimeEq returns 1 if x == y and 0 otherwise.
var ConstantTimeEq = subtle.ConstantTimeEq

// ConstantTimeLessOrEq returns 1 if x &lt;= y and 0 otherwise.
// Its behavior is undefined if x or y are negative or &gt; 2**31 - 1.
var ConstantTimeLessOrEq = subtle.ConstantTimeLessOrEq

// ConstantTimeSelect returns x if v is 1 and y if v is 0.
// Its behavior is undefined if v takes any other value.
var ConstantTimeSelect = subtle.ConstantTimeSelect
