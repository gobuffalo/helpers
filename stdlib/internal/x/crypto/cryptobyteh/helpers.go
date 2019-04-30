package cryptobyteh

import (
	"internal/x/crypto/cryptobyte"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewBuilderKey = "NewBuilder"

	NewFixedBuilderKey = "NewFixedBuilder"
)

func New() hctx.Map {
	return hctx.Map{

		NewBuilderKey: NewBuilder,

		NewFixedBuilderKey: NewFixedBuilder,
	}
}

// NewBuilder creates a Builder that appends its output to the given buffer.
// Like append(), the slice will be reallocated if its capacity is exceeded.
// Use Bytes to get the final buffer.
var NewBuilder = cryptobyte.NewBuilder

// NewFixedBuilder creates a Builder that appends its output into the given
// buffer. This builder does not reallocate the output buffer. Writes that
// would exceed the buffer&#39;s capacity are treated as an error.
var NewFixedBuilder = cryptobyte.NewFixedBuilder
