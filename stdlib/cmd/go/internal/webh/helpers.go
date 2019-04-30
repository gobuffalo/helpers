package webh

import (
	"cmd/go/internal/web"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	GetKey = "Get"

	GetKey = "Get"

	GetMaybeInsecureKey = "GetMaybeInsecure"

	GetMaybeInsecureKey = "GetMaybeInsecure"

	OpenBrowserKey = "OpenBrowser"

	OpenBrowserKey = "OpenBrowser"

	QueryEscapeKey = "QueryEscape"

	QueryEscapeKey = "QueryEscape"
)

func New() hctx.Map {
	return hctx.Map{

		GetKey: Get,

		GetKey: Get,

		GetMaybeInsecureKey: GetMaybeInsecure,

		GetMaybeInsecureKey: GetMaybeInsecure,

		OpenBrowserKey: OpenBrowser,

		OpenBrowserKey: OpenBrowser,

		QueryEscapeKey: QueryEscape,

		QueryEscapeKey: QueryEscape,
	}
}

var Get = web.Get

// Get returns the data from an HTTP GET request for the given URL.
var Get = web.Get

var GetMaybeInsecure = web.GetMaybeInsecure

// GetMaybeInsecure returns the body of either the importPath&#39;s
// https resource or, if unavailable and permitted by the security mode, the http resource.
var GetMaybeInsecure = web.GetMaybeInsecure

var OpenBrowser = web.OpenBrowser

var OpenBrowser = web.OpenBrowser

var QueryEscape = web.QueryEscape

var QueryEscape = web.QueryEscape
