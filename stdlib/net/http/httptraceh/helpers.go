package httptraceh

import (
	"net/http/httptrace"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ContextClientTraceKey = "ContextClientTrace"

	WithClientTraceKey = "WithClientTrace"
)

func New() hctx.Map {
	return hctx.Map{

		ContextClientTraceKey: ContextClientTrace,

		WithClientTraceKey: WithClientTrace,
	}
}

// ContextClientTrace returns the ClientTrace associated with the
// provided context. If none, it returns nil.
var ContextClientTrace = httptrace.ContextClientTrace

// WithClientTrace returns a new context based on the provided parent
// ctx. HTTP client requests made with the returned context will use
// the provided trace hooks, in addition to any previous hooks
// registered with ctx. Any hooks defined in the provided trace will
// be called first.
var WithClientTrace = httptrace.WithClientTrace
