package bidiruleh

import (
	"internal/x/text/secure/bidirule"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DirectionKey = "Direction"

	DirectionStringKey = "DirectionString"

	NewKey = "New"

	ValidKey = "Valid"

	ValidStringKey = "ValidString"
)

func New() hctx.Map {
	return hctx.Map{

		DirectionKey: Direction,

		DirectionStringKey: DirectionString,

		NewKey: New,

		ValidKey: Valid,

		ValidStringKey: ValidString,
	}
}

// Direction reports the direction of the given label as defined by RFC 5893.
// The Bidi Rule does not have to be applied to labels of the category
// LeftToRight.
var Direction = bidirule.Direction

// DirectionString reports the direction of the given label as defined by RFC
// 5893. The Bidi Rule does not have to be applied to labels of the category
// LeftToRight.
var DirectionString = bidirule.DirectionString

// New returns a Transformer that verifies that input adheres to the Bidi Rule.
var New = bidirule.New

// Valid reports whether b conforms to the BiDi rule.
var Valid = bidirule.Valid

// ValidString reports whether s conforms to the BiDi rule.
var ValidString = bidirule.ValidString
