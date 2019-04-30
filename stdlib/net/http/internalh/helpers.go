package internalh

import (
	"net/http/internal"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewChunkedReaderKey = "NewChunkedReader"

	NewChunkedWriterKey = "NewChunkedWriter"
)

func New() hctx.Map {
	return hctx.Map{

		NewChunkedReaderKey: NewChunkedReader,

		NewChunkedWriterKey: NewChunkedWriter,
	}
}

// NewChunkedReader returns a new chunkedReader that translates the data read from r
// out of HTTP &#34;chunked&#34; format before returning it.
// The chunkedReader returns io.EOF when the final 0-length chunk is read.
//
// NewChunkedReader is not needed by normal applications. The http package
// automatically decodes chunking when reading response bodies.
var NewChunkedReader = internal.NewChunkedReader

// NewChunkedWriter returns a new chunkedWriter that translates writes into HTTP
// &#34;chunked&#34; format before writing them to w. Closing the returned chunkedWriter
// sends the final 0-length chunk that marks the end of the stream but does
// not send the final CRLF that appears after trailers; trailers and the last
// CRLF must be written separately.
//
// NewChunkedWriter is not needed by normal applications. The http
// package adds chunking automatically if handlers don&#39;t set a
// Content-Length header. Using newChunkedWriter inside a handler
// would result in double chunking or chunking with a Content-Length
// length, both of which are wrong.
var NewChunkedWriter = internal.NewChunkedWriter
