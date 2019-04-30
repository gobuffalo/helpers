package proftesth

import (
	"cmd/vendor/github.com/google/pprof/internal/proftest"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	DiffKey = "Diff"

	EncodeJSONKey = "EncodeJSON"
)

func New() hctx.Map {
	return hctx.Map{

		DiffKey: Diff,

		EncodeJSONKey: EncodeJSON,
	}
}

// Diff compares two byte arrays using the diff tool to highlight the
// differences. It is meant for testing purposes to display the
// differences between expected and actual output.
var Diff = proftest.Diff

// EncodeJSON encodes a value into a byte array. This is intended for
// testing purposes.
var EncodeJSON = proftest.EncodeJSON
