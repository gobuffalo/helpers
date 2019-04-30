package drawh

import (
	"image/draw"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DrawKey = "Draw"

	DrawMaskKey = "DrawMask"
)

func New() hctx.Map {
	return hctx.Map{

		DrawKey: Draw,

		DrawMaskKey: DrawMask,
	}
}

// Draw calls DrawMask with a nil mask.
var Draw = draw.Draw

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result of a Porter-Duff composition. A nil mask is treated as opaque.
var DrawMask = draw.DrawMask
