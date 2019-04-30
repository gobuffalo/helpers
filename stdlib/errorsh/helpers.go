package errorsh

import (
	"errors"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewKey = "New"
)

func New() hctx.Map {
	return hctx.Map{

		NewKey: New,
	}
}

// New returns an error that formats as the given text.
var New = errors.New
