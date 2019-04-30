package utf8h

import (
	"unicode/utf8"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DecodeLastRuneKey = "DecodeLastRune"

	DecodeLastRuneInStringKey = "DecodeLastRuneInString"

	DecodeRuneKey = "DecodeRune"

	DecodeRuneInStringKey = "DecodeRuneInString"

	EncodeRuneKey = "EncodeRune"

	FullRuneKey = "FullRune"

	FullRuneInStringKey = "FullRuneInString"

	RuneCountKey = "RuneCount"

	RuneCountInStringKey = "RuneCountInString"

	RuneLenKey = "RuneLen"

	RuneStartKey = "RuneStart"

	ValidKey = "Valid"

	ValidRuneKey = "ValidRune"

	ValidStringKey = "ValidString"
)

func New() hctx.Map {
	return hctx.Map{

		DecodeLastRuneKey: DecodeLastRune,

		DecodeLastRuneInStringKey: DecodeLastRuneInString,

		DecodeRuneKey: DecodeRune,

		DecodeRuneInStringKey: DecodeRuneInString,

		EncodeRuneKey: EncodeRune,

		FullRuneKey: FullRune,

		FullRuneInStringKey: FullRuneInString,

		RuneCountKey: RuneCount,

		RuneCountInStringKey: RuneCountInString,

		RuneLenKey: RuneLen,

		RuneStartKey: RuneStart,

		ValidKey: Valid,

		ValidRuneKey: ValidRune,

		ValidStringKey: ValidString,
	}
}

// DecodeLastRune unpacks the last UTF-8 encoding in p and returns the rune and
// its width in bytes. If p is empty it returns (RuneError, 0). Otherwise, if
// the encoding is invalid, it returns (RuneError, 1). Both are impossible
// results for correct, non-empty UTF-8.
//
// An encoding is invalid if it is incorrect UTF-8, encodes a rune that is
// out of range, or is not the shortest possible UTF-8 encoding for the
// value. No other validation is performed.
var DecodeLastRune = utf8.DecodeLastRune

// DecodeLastRuneInString is like DecodeLastRune but its input is a string. If
// s is empty it returns (RuneError, 0). Otherwise, if the encoding is invalid,
// it returns (RuneError, 1). Both are impossible results for correct,
// non-empty UTF-8.
//
// An encoding is invalid if it is incorrect UTF-8, encodes a rune that is
// out of range, or is not the shortest possible UTF-8 encoding for the
// value. No other validation is performed.
var DecodeLastRuneInString = utf8.DecodeLastRuneInString

// DecodeRune unpacks the first UTF-8 encoding in p and returns the rune and
// its width in bytes. If p is empty it returns (RuneError, 0). Otherwise, if
// the encoding is invalid, it returns (RuneError, 1). Both are impossible
// results for correct, non-empty UTF-8.
//
// An encoding is invalid if it is incorrect UTF-8, encodes a rune that is
// out of range, or is not the shortest possible UTF-8 encoding for the
// value. No other validation is performed.
var DecodeRune = utf8.DecodeRune

// DecodeRuneInString is like DecodeRune but its input is a string. If s is
// empty it returns (RuneError, 0). Otherwise, if the encoding is invalid, it
// returns (RuneError, 1). Both are impossible results for correct, non-empty
// UTF-8.
//
// An encoding is invalid if it is incorrect UTF-8, encodes a rune that is
// out of range, or is not the shortest possible UTF-8 encoding for the
// value. No other validation is performed.
var DecodeRuneInString = utf8.DecodeRuneInString

// EncodeRune writes into p (which must be large enough) the UTF-8 encoding of the rune.
// It returns the number of bytes written.
var EncodeRune = utf8.EncodeRune

// FullRune reports whether the bytes in p begin with a full UTF-8 encoding of a rune.
// An invalid encoding is considered a full Rune since it will convert as a width-1 error rune.
var FullRune = utf8.FullRune

// FullRuneInString is like FullRune but its input is a string.
var FullRuneInString = utf8.FullRuneInString

// RuneCount returns the number of runes in p. Erroneous and short
// encodings are treated as single runes of width 1 byte.
var RuneCount = utf8.RuneCount

// RuneCountInString is like RuneCount but its input is a string.
var RuneCountInString = utf8.RuneCountInString

// RuneLen returns the number of bytes required to encode the rune.
// It returns -1 if the rune is not a valid value to encode in UTF-8.
var RuneLen = utf8.RuneLen

// RuneStart reports whether the byte could be the first byte of an encoded,
// possibly invalid rune. Second and subsequent bytes always have the top two
// bits set to 10.
var RuneStart = utf8.RuneStart

// Valid reports whether p consists entirely of valid UTF-8-encoded runes.
var Valid = utf8.Valid

// ValidRune reports whether r can be legally encoded as UTF-8.
// Code points that are out of range or a surrogate half are illegal.
var ValidRune = utf8.ValidRune

// ValidString reports whether s consists entirely of valid UTF-8-encoded runes.
var ValidString = utf8.ValidString
