package jpegh

import (
	"image/jpeg"

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

// Decode reads a JPEG image from r and returns it as an image.Image.
var Decode = jpeg.Decode

// DecodeConfig returns the color model and dimensions of a JPEG image without
// decoding the entire image.
var DecodeConfig = jpeg.DecodeConfig

// Encode writes the Image m to w in JPEG 4:2:0 baseline format with the given
// options. Default parameters are used if a nil *Options is passed.
var Encode = jpeg.Encode
