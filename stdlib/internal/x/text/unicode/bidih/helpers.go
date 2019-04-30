package bidih

import (
	"internal/x/text/unicode/bidi"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AppendReverseKey = "AppendReverse"

	DefaultDirectionKey = "DefaultDirection"

	LookupKey = "Lookup"

	LookupRuneKey = "LookupRune"

	LookupStringKey = "LookupString"

	ReverseStringKey = "ReverseString"
)

func New() hctx.Map {
	return hctx.Map{

		AppendReverseKey: AppendReverse,

		DefaultDirectionKey: DefaultDirection,

		LookupKey: Lookup,

		LookupRuneKey: LookupRune,

		LookupStringKey: LookupString,

		ReverseStringKey: ReverseString,
	}
}

// AppendReverse reverses the order of characters of in, appends them to out,
// and returns the result. Modifiers will still follow the runes they modify.
// Brackets are replaced with their counterparts.
var AppendReverse = bidi.AppendReverse

// DefaultDirection sets the default direction for a Paragraph. The direction is
// overridden if the text contains directional characters.
var DefaultDirection = bidi.DefaultDirection

// Lookup returns properties for the first rune in s and the width in bytes of
// its encoding. The size will be 0 if s does not hold enough bytes to complete
// the encoding.
var Lookup = bidi.Lookup

// LookupRune returns properties for r.
var LookupRune = bidi.LookupRune

// LookupString returns properties for the first rune in s and the width in
// bytes of its encoding. The size will be 0 if s does not hold enough bytes to
// complete the encoding.
var LookupString = bidi.LookupString

// ReverseString reverses the order of characters in s and returns a new string.
// Modifiers will still follow the runes they modify. Brackets are replaced with
// their counterparts.
var ReverseString = bidi.ReverseString
