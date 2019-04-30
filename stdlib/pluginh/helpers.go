package pluginh

import (
	"plugin"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	OpenKey = "Open"
)

func New() hctx.Map {
	return hctx.Map{

		OpenKey: Open,
	}
}

// Open opens a Go plugin.
// If a path has already been opened, then the existing *Plugin is returned.
// It is safe for concurrent use by multiple goroutines.
var Open = plugin.Open
