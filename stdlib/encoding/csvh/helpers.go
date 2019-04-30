package csvh

import (
	"encoding/csv"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewReaderKey = "NewReader"

	NewWriterKey = "NewWriter"
)

func New() hctx.Map {
	return hctx.Map{

		NewReaderKey: NewReader,

		NewWriterKey: NewWriter,
	}
}

// NewReader returns a new Reader that reads from r.
var NewReader = csv.NewReader

// NewWriter returns a new Writer that writes to w.
var NewWriter = csv.NewWriter
