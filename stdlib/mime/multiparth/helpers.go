package multiparth

import (
	"mime/multipart"

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

// NewReader creates a new multipart Reader reading from r using the
// given MIME boundary.
//
// The boundary is usually obtained from the &#34;boundary&#34; parameter of
// the message&#39;s &#34;Content-Type&#34; header. Use mime.ParseMediaType to
// parse such headers.
var NewReader = multipart.NewReader

// NewWriter returns a new multipart Writer with a random boundary,
// writing to w.
var NewWriter = multipart.NewWriter
