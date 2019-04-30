package routeh

import (
	"internal/x/net/route"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	FetchRIBKey = "FetchRIB"

	ParseRIBKey = "ParseRIB"
)

func New() hctx.Map {
	return hctx.Map{

		FetchRIBKey: FetchRIB,

		ParseRIBKey: ParseRIB,
	}
}

// FetchRIB fetches a routing information base from the operating
// system.
//
// The provided af must be an address family.
//
// The provided arg must be a RIBType-specific argument.
// When RIBType is related to routes, arg might be a set of route
// flags. When RIBType is related to network interfaces, arg might be
// an interface index or a set of interface flags. In most cases, zero
// means a wildcard.
var FetchRIB = route.FetchRIB

// ParseRIB parses b as a routing information base and returns a list
// of routing messages.
var ParseRIB = route.ParseRIB
