package httputilh

import (
	"net/http/httputil"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DumpRequestKey = "DumpRequest"

	DumpRequestOutKey = "DumpRequestOut"

	DumpResponseKey = "DumpResponse"

	ErrLineTooLongKey = "ErrLineTooLong"

	NewChunkedReaderKey = "NewChunkedReader"

	NewChunkedWriterKey = "NewChunkedWriter"

	NewClientConnKey = "NewClientConn"

	NewProxyClientConnKey = "NewProxyClientConn"

	NewServerConnKey = "NewServerConn"

	NewSingleHostReverseProxyKey = "NewSingleHostReverseProxy"
)

func New() hctx.Map {
	return hctx.Map{

		DumpRequestKey: DumpRequest,

		DumpRequestOutKey: DumpRequestOut,

		DumpResponseKey: DumpResponse,

		ErrLineTooLongKey: ErrLineTooLong,

		NewChunkedReaderKey: NewChunkedReader,

		NewChunkedWriterKey: NewChunkedWriter,

		NewClientConnKey: NewClientConn,

		NewProxyClientConnKey: NewProxyClientConn,

		NewServerConnKey: NewServerConn,

		NewSingleHostReverseProxyKey: NewSingleHostReverseProxy,
	}
}

// DumpRequest returns the given request in its HTTP/1.x wire
// representation. It should only be used by servers to debug client
// requests. The returned representation is an approximation only;
// some details of the initial request are lost while parsing it into
// an http.Request. In particular, the order and case of header field
// names are lost. The order of values in multi-valued headers is kept
// intact. HTTP/2 requests are dumped in HTTP/1.x form, not in their
// original binary representations.
//
// If body is true, DumpRequest also returns the body. To do so, it
// consumes req.Body and then replaces it with a new io.ReadCloser
// that yields the same bytes. If DumpRequest returns an error,
// the state of req is undefined.
//
// The documentation for http.Request.Write details which fields
// of req are included in the dump.
var DumpRequest = httputil.DumpRequest

// DumpRequestOut is like DumpRequest but for outgoing client requests. It
// includes any headers that the standard http.Transport adds, such as
// User-Agent.
var DumpRequestOut = httputil.DumpRequestOut

// DumpResponse is like DumpRequest but dumps a response.
var DumpResponse = httputil.DumpResponse

var ErrLineTooLong = httputil.ErrLineTooLong

// NewChunkedReader returns a new chunkedReader that translates the data read from r
// out of HTTP &#34;chunked&#34; format before returning it.
// The chunkedReader returns io.EOF when the final 0-length chunk is read.
//
// NewChunkedReader is not needed by normal applications. The http package
// automatically decodes chunking when reading response bodies.
var NewChunkedReader = httputil.NewChunkedReader

// NewChunkedWriter returns a new chunkedWriter that translates writes into HTTP
// &#34;chunked&#34; format before writing them to w. Closing the returned chunkedWriter
// sends the final 0-length chunk that marks the end of the stream but does
// not send the final CRLF that appears after trailers; trailers and the last
// CRLF must be written separately.
//
// NewChunkedWriter is not needed by normal applications. The http
// package adds chunking automatically if handlers don&#39;t set a
// Content-Length header. Using NewChunkedWriter inside a handler
// would result in double chunking or chunking with a Content-Length
// length, both of which are wrong.
var NewChunkedWriter = httputil.NewChunkedWriter

// NewClientConn is an artifact of Go&#39;s early HTTP implementation.
// It is low-level, old, and unused by Go&#39;s current HTTP stack.
// We should have deleted it before Go 1.
//
// Deprecated: Use the Client or Transport in package net/http instead.
var NewClientConn = httputil.NewClientConn

// NewProxyClientConn is an artifact of Go&#39;s early HTTP implementation.
// It is low-level, old, and unused by Go&#39;s current HTTP stack.
// We should have deleted it before Go 1.
//
// Deprecated: Use the Client or Transport in package net/http instead.
var NewProxyClientConn = httputil.NewProxyClientConn

// NewServerConn is an artifact of Go&#39;s early HTTP implementation.
// It is low-level, old, and unused by Go&#39;s current HTTP stack.
// We should have deleted it before Go 1.
//
// Deprecated: Use the Server in package net/http instead.
var NewServerConn = httputil.NewServerConn

// NewSingleHostReverseProxy returns a new ReverseProxy that routes
// URLs to the scheme, host, and base path provided in target. If the
// target&#39;s path is &#34;/base&#34; and the incoming request was for &#34;/dir&#34;,
// the target request will be for /base/dir.
// NewSingleHostReverseProxy does not rewrite the Host header.
// To rewrite Host headers, use ReverseProxy directly with a custom
// Director policy.
var NewSingleHostReverseProxy = httputil.NewSingleHostReverseProxy
