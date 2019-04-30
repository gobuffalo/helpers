package fcgih

import (
	"net/http/fcgi"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ProcessEnvKey = "ProcessEnv"

	ServeKey = "Serve"
)

func New() hctx.Map {
	return hctx.Map{

		ProcessEnvKey: ProcessEnv,

		ServeKey: Serve,
	}
}

// ProcessEnv returns FastCGI environment variables associated with the request r
// for which no effort was made to be included in the request itself - the data
// is hidden in the request&#39;s context. As an example, if REMOTE_USER is set for a
// request, it will not be found anywhere in r, but it will be included in
// ProcessEnv&#39;s response (via r&#39;s context).
var ProcessEnv = fcgi.ProcessEnv

// Serve accepts incoming FastCGI connections on the listener l, creating a new
// goroutine for each. The goroutine reads requests and then calls handler
// to reply to them.
// If l is nil, Serve accepts connections from os.Stdin.
// If handler is nil, http.DefaultServeMux is used.
var Serve = fcgi.Serve
