package hmach

import (
	"crypto/hmac"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	EqualKey = "Equal"
)

func New() hctx.Map {
	return hctx.Map{

		EqualKey: Equal,
	}
}

// Equal compares two MACs for equality without leaking timing information.
var Equal = hmac.Equal
