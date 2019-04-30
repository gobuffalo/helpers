package edith

import (
	"cmd/internal/edit"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewBufferKey = "NewBuffer"
)

func New() hctx.Map {
	return hctx.Map{

		NewBufferKey: NewBuffer,
	}
}

// NewBuffer returns a new buffer to accumulate changes to an initial data slice.
// The returned buffer maintains a reference to the data, so the caller must ensure
// the data is not modified until after the Buffer is done being used.
var NewBuffer = edit.NewBuffer
