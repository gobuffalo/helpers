package dnsmessageh

import (
	"internal/x/net/dns/dnsmessage"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewBuilderKey = "NewBuilder"

	NewNameKey = "NewName"
)

func New() hctx.Map {
	return hctx.Map{

		NewBuilderKey: NewBuilder,

		NewNameKey: NewName,
	}
}

// NewBuilder creates a new builder with compression disabled.
//
// Note: Most users will want to immediately enable compression with the
// EnableCompression method. See that method&#39;s comment for why you may or may
// not want to enable compression.
//
// The DNS message is appended to the provided initial buffer buf (which may be
// nil) as it is built. The final message is returned by the (*Builder).Finish
// method, which may return the same underlying array if there was sufficient
// capacity in the slice.
var NewBuilder = dnsmessage.NewBuilder

// NewName creates a new Name from a string.
var NewName = dnsmessage.NewName
