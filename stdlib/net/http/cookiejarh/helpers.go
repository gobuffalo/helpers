package cookiejarh

import (
	"net/http/cookiejar"

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

// New returns a new cookie jar. A nil *Options is equivalent to a zero
// Options.
var New = cookiejar.New
