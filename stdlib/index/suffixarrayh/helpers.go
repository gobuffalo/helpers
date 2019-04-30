package suffixarrayh

import (
	"index/suffixarray"

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

// New creates a new Index for data.
// Index creation time is O(N*log(N)) for N = len(data).
var New = suffixarray.New
