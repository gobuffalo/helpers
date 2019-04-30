package transporth

import (
	"cmd/vendor/github.com/google/pprof/internal/transport"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewKey = "New"
)

func New() hctx.Map {
	return hctx.Map{

		NewKey: New,
	}
}

// New returns a round tripper for making requests with the
// specified cert, key, and ca. The flags tls_cert, tls_key, and tls_ca are
// added to the flagset to allow a user to specify the cert, key, and ca. If
// the flagset is nil, no flags will be added, and users will not be able to
// use these flags.
var New = transport.New
