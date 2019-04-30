package colorh

import (
	"image/color"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	CMYKToRGBKey = "CMYKToRGB"

	ModelFuncKey = "ModelFunc"

	RGBToCMYKKey = "RGBToCMYK"

	RGBToYCbCrKey = "RGBToYCbCr"

	YCbCrToRGBKey = "YCbCrToRGB"
)

func New() hctx.Map {
	return hctx.Map{

		CMYKToRGBKey: CMYKToRGB,

		ModelFuncKey: ModelFunc,

		RGBToCMYKKey: RGBToCMYK,

		RGBToYCbCrKey: RGBToYCbCr,

		YCbCrToRGBKey: YCbCrToRGB,
	}
}

// CMYKToRGB converts a CMYK quadruple to an RGB triple.
var CMYKToRGB = color.CMYKToRGB

// ModelFunc returns a Model that invokes f to implement the conversion.
var ModelFunc = color.ModelFunc

// RGBToCMYK converts an RGB triple to a CMYK quadruple.
var RGBToCMYK = color.RGBToCMYK

// RGBToYCbCr converts an RGB triple to a Y&#39;CbCr triple.
var RGBToYCbCr = color.RGBToYCbCr

// YCbCrToRGB converts a Y&#39;CbCr triple to an RGB triple.
var YCbCrToRGB = color.YCbCrToRGB
