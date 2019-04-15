package paths

import "github.com/gobuffalo/helpers/hctx"

func New() hctx.Map {
	return hctx.Map{
		"assetPath": Asset,
	}
}
