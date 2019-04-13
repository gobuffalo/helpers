package encoders

import "github.com/gobuffalo/helpers/hctx"

func New() hctx.Map {
	return hctx.Map{
		"json":   ToJSON,
		"raw":    Raw,
		"toJSON": ToJSON,
	}
}
