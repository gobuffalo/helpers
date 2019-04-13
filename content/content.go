package content

import "github.com/gobuffalo/helpers/hctx"

func Map() hctx.Map {
	return hctx.Map{
		"contentOf":  ContentOf,
		"contentFor": ContentFor,
	}
}
