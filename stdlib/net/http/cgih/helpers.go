package cgih

import (
	"net/http/cgi"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	RequestKey = "Request"

	RequestFromMapKey = "RequestFromMap"

	ServeKey = "Serve"
)

func New() hctx.Map {
	return hctx.Map{

		RequestKey: Request,

		RequestFromMapKey: RequestFromMap,

		ServeKey: Serve,
	}
}

// Request returns the HTTP request as represented in the current
// environment. This assumes the current program is being run
// by a web server in a CGI environment.
// The returned Request&#39;s Body is populated, if applicable.
var Request = cgi.Request

// RequestFromMap creates an http.Request from CGI variables.
// The returned Request&#39;s Body field is not populated.
var RequestFromMap = cgi.RequestFromMap

// Serve executes the provided Handler on the currently active CGI
// request, if any. If there&#39;s no current CGI environment
// an error is returned. The provided handler may be nil to use
// http.DefaultServeMux.
var Serve = cgi.Serve
