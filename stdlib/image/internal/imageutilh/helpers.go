package imageutilh

import (
	"image/internal/imageutil"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DrawYCbCrKey = "DrawYCbCr"
)

func New() hctx.Map {
	return hctx.Map{

		DrawYCbCrKey: DrawYCbCr,
	}
}

// DrawYCbCr draws the YCbCr source image on the RGBA destination image with
// r.Min in dst aligned with sp in src. It reports whether the draw was
// successful. If it returns false, no dst pixels were changed.
//
// This function assumes that r is entirely within dst&#39;s bounds and the
// translation of r from dst coordinate space to src coordinate space is
// entirely within src&#39;s bounds.
var DrawYCbCr = imageutil.DrawYCbCr
