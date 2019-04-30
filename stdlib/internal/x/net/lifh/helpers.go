package lifh

import (
	"internal/x/net/lif"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AddrsKey = "Addrs"

	LinksKey = "Links"
)

func New() hctx.Map {
	return hctx.Map{

		AddrsKey: Addrs,

		LinksKey: Links,
	}
}

// Addrs returns a list of interface addresses.
//
// The provided af must be an address family and name must be a data
// link name. The zero value of af or name means a wildcard.
var Addrs = lif.Addrs

// Links returns a list of logical data links.
//
// The provided af must be an address family and name must be a data
// link name. The zero value of af or name means a wildcard.
var Links = lif.Links
