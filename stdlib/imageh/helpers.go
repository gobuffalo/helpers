package imageh

import (
	"image"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DecodeKey = "Decode"

	DecodeConfigKey = "DecodeConfig"

	NewAlphaKey = "NewAlpha"

	NewAlpha16Key = "NewAlpha16"

	NewCMYKKey = "NewCMYK"

	NewGrayKey = "NewGray"

	NewGray16Key = "NewGray16"

	NewNRGBAKey = "NewNRGBA"

	NewNRGBA64Key = "NewNRGBA64"

	NewNYCbCrAKey = "NewNYCbCrA"

	NewPalettedKey = "NewPaletted"

	NewRGBAKey = "NewRGBA"

	NewRGBA64Key = "NewRGBA64"

	NewUniformKey = "NewUniform"

	NewYCbCrKey = "NewYCbCr"

	PtKey = "Pt"

	RectKey = "Rect"

	RegisterFormatKey = "RegisterFormat"
)

func New() hctx.Map {
	return hctx.Map{

		DecodeKey: Decode,

		DecodeConfigKey: DecodeConfig,

		NewAlphaKey: NewAlpha,

		NewAlpha16Key: NewAlpha16,

		NewCMYKKey: NewCMYK,

		NewGrayKey: NewGray,

		NewGray16Key: NewGray16,

		NewNRGBAKey: NewNRGBA,

		NewNRGBA64Key: NewNRGBA64,

		NewNYCbCrAKey: NewNYCbCrA,

		NewPalettedKey: NewPaletted,

		NewRGBAKey: NewRGBA,

		NewRGBA64Key: NewRGBA64,

		NewUniformKey: NewUniform,

		NewYCbCrKey: NewYCbCr,

		PtKey: Pt,

		RectKey: Rect,

		RegisterFormatKey: RegisterFormat,
	}
}

// Decode decodes an image that has been encoded in a registered format.
// The string returned is the format name used during format registration.
// Format registration is typically done by an init function in the codec-
// specific package.
var Decode = image.Decode

// DecodeConfig decodes the color model and dimensions of an image that has
// been encoded in a registered format. The string returned is the format name
// used during format registration. Format registration is typically done by
// an init function in the codec-specific package.
var DecodeConfig = image.DecodeConfig

// NewAlpha returns a new Alpha image with the given bounds.
var NewAlpha = image.NewAlpha

// NewAlpha16 returns a new Alpha16 image with the given bounds.
var NewAlpha16 = image.NewAlpha16

// NewCMYK returns a new CMYK image with the given bounds.
var NewCMYK = image.NewCMYK

// NewGray returns a new Gray image with the given bounds.
var NewGray = image.NewGray

// NewGray16 returns a new Gray16 image with the given bounds.
var NewGray16 = image.NewGray16

// NewNRGBA returns a new NRGBA image with the given bounds.
var NewNRGBA = image.NewNRGBA

// NewNRGBA64 returns a new NRGBA64 image with the given bounds.
var NewNRGBA64 = image.NewNRGBA64

// NewNYCbCrA returns a new NYCbCrA image with the given bounds and subsample
// ratio.
var NewNYCbCrA = image.NewNYCbCrA

// NewPaletted returns a new Paletted image with the given width, height and
// palette.
var NewPaletted = image.NewPaletted

// NewRGBA returns a new RGBA image with the given bounds.
var NewRGBA = image.NewRGBA

// NewRGBA64 returns a new RGBA64 image with the given bounds.
var NewRGBA64 = image.NewRGBA64

var NewUniform = image.NewUniform

// NewYCbCr returns a new YCbCr image with the given bounds and subsample
// ratio.
var NewYCbCr = image.NewYCbCr

// Pt is shorthand for Point{X, Y}.
var Pt = image.Pt

// Rect is shorthand for Rectangle{Pt(x0, y0), Pt(x1, y1)}. The returned
// rectangle has minimum and maximum coordinates swapped if necessary so that
// it is well-formed.
var Rect = image.Rect

// RegisterFormat registers an image format for use by Decode.
// Name is the name of the format, like &#34;jpeg&#34; or &#34;png&#34;.
// Magic is the magic prefix that identifies the format&#39;s encoding. The magic
// string can contain &#34;?&#34; wildcards that each match any one byte.
// Decode is the function that decodes the encoded image.
// DecodeConfig is the function that decodes just its configuration.
var RegisterFormat = image.RegisterFormat
