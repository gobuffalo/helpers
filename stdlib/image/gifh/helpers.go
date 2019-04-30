package gifh

import (
	"image/gif"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DecodeKey = "Decode"

	DecodeAllKey = "DecodeAll"

	DecodeConfigKey = "DecodeConfig"

	EncodeKey = "Encode"

	EncodeAllKey = "EncodeAll"
)

func New() hctx.Map {
	return hctx.Map{

		DecodeKey: Decode,

		DecodeAllKey: DecodeAll,

		DecodeConfigKey: DecodeConfig,

		EncodeKey: Encode,

		EncodeAllKey: EncodeAll,
	}
}

// Decode reads a GIF image from r and returns the first embedded
// image as an image.Image.
var Decode = gif.Decode

// DecodeAll reads a GIF image from r and returns the sequential frames
// and timing information.
var DecodeAll = gif.DecodeAll

// DecodeConfig returns the global color model and dimensions of a GIF image
// without decoding the entire image.
var DecodeConfig = gif.DecodeConfig

// Encode writes the Image m to w in GIF format.
var Encode = gif.Encode

// EncodeAll writes the images in g to w in GIF format with the
// given loop count and delay between frames.
var EncodeAll = gif.EncodeAll
