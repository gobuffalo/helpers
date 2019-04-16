package content

import "github.com/gobuffalo/helpers/hctx"

const (
	OfKey  = "contentOf"
	ForKey = "contentFor"
)

func New() hctx.Map {
	return hctx.Map{
		OfKey:  ContentOf,
		ForKey: ContentFor,
	}
}
