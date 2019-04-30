package tabwriterh

import (
	"text/tabwriter"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewWriterKey = "NewWriter"
)

func New() hctx.Map {
	return hctx.Map{

		NewWriterKey: NewWriter,
	}
}

// NewWriter allocates and initializes a new tabwriter.Writer.
// The parameters are the same as for the Init function.
var NewWriter = tabwriter.NewWriter
