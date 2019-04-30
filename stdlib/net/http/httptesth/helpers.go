package httptesth

import (
	"net/http/httptest"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewRecorderKey = "NewRecorder"

	NewRequestKey = "NewRequest"

	NewServerKey = "NewServer"

	NewTLSServerKey = "NewTLSServer"

	NewUnstartedServerKey = "NewUnstartedServer"
)

func New() hctx.Map {
	return hctx.Map{

		NewRecorderKey: NewRecorder,

		NewRequestKey: NewRequest,

		NewServerKey: NewServer,

		NewTLSServerKey: NewTLSServer,

		NewUnstartedServerKey: NewUnstartedServer,
	}
}

// NewRecorder returns an initialized ResponseRecorder.
var NewRecorder = httptest.NewRecorder

// NewRequest returns a new incoming server Request, suitable
// for passing to an http.Handler for testing.
//
// The target is the RFC 7230 &#34;request-target&#34;: it may be either a
// path or an absolute URL. If target is an absolute URL, the host name
// from the URL is used. Otherwise, &#34;example.com&#34; is used.
//
// The TLS field is set to a non-nil dummy value if target has scheme
// &#34;https&#34;.
//
// The Request.Proto is always HTTP/1.1.
//
// An empty method means &#34;GET&#34;.
//
// The provided body may be nil. If the body is of type *bytes.Reader,
// *strings.Reader, or *bytes.Buffer, the Request.ContentLength is
// set.
//
// NewRequest panics on error for ease of use in testing, where a
// panic is acceptable.
//
// To generate a client HTTP request instead of a server request, see
// the NewRequest function in the net/http package.
var NewRequest = httptest.NewRequest

// NewServer starts and returns a new Server.
// The caller should call Close when finished, to shut it down.
var NewServer = httptest.NewServer

// NewTLSServer starts and returns a new Server using TLS.
// The caller should call Close when finished, to shut it down.
var NewTLSServer = httptest.NewTLSServer

// NewUnstartedServer returns a new Server but doesn&#39;t start it.
//
// After changing its configuration, the caller should call Start or
// StartTLS.
//
// The caller should call Close when finished, to shut it down.
var NewUnstartedServer = httptest.NewUnstartedServer
