package httpproxyh

import (
	"internal/x/net/http/httpproxy"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	FromEnvironmentKey = "FromEnvironment"
)

func New() hctx.Map {
	return hctx.Map{

		FromEnvironmentKey: FromEnvironment,
	}
}

// FromEnvironment returns a Config instance populated from the
// environment variables HTTP_PROXY, HTTPS_PROXY and NO_PROXY (or the
// lowercase versions thereof). HTTPS_PROXY takes precedence over
// HTTP_PROXY for https requests.
//
// The environment values may be either a complete URL or a
// &#34;host[:port]&#34;, in which case the &#34;http&#34; scheme is assumed. An error
// is returned if the value is a different form.
var FromEnvironment = httpproxy.FromEnvironment
