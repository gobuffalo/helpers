package curve25519h

import (
	"internal/x/crypto/curve25519"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ScalarBaseMultKey = "ScalarBaseMult"

	ScalarMultKey = "ScalarMult"
)

func New() hctx.Map {
	return hctx.Map{

		ScalarBaseMultKey: ScalarBaseMult,

		ScalarMultKey: ScalarMult,
	}
}

// ScalarBaseMult sets dst to the product in*base where dst and base are the x
// coordinates of group points, base is the standard generator and all values
// are in little-endian form.
var ScalarBaseMult = curve25519.ScalarBaseMult

// ScalarMult sets dst to the product in*base where dst and base are the x
// coordinates of group points and all values are in little-endian form.
var ScalarMult = curve25519.ScalarMult
