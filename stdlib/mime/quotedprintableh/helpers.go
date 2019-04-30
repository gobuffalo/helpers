package quotedprintableh

import (
	"mime/quotedprintable"

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

// NewReader returns a quoted-printable reader, decoding from r.
var NewReader = quotedprintable.NewReader

// NewWriter returns a new Writer that writes to w.
var NewWriter = quotedprintable.NewWriter
