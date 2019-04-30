package pngh

import (
	"image/png"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DecodeKey = "Decode"

	DecodeConfigKey = "DecodeConfig"

	EncodeKey = "Encode"
)

func New() hctx.Map {
	return hctx.Map{

		DecodeKey: Decode,

		DecodeConfigKey: DecodeConfig,

		EncodeKey: Encode,
	}
}

// Decode reads a PNG image from r and returns it as an image.Image.
// The type of Image returned depends on the PNG contents.
var Decode = png.Decode

// DecodeConfig returns the color model and dimensions of a PNG image without
// decoding the entire image.
var DecodeConfig = png.DecodeConfig

// Encode writes the Image m to w in PNG format. Any Image may be
// encoded, but images that are not image.NRGBA might be encoded lossily.
var Encode = png.Encode
