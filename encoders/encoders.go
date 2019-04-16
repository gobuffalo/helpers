package encoders

import "github.com/gobuffalo/helpers/hctx"

const (
	ToJSONKey = "toJSON"
	RawKey    = "raw"
)

func New() hctx.Map {
	return hctx.Map{
		"json":    ToJSON,
		RawKey:    Raw,
		ToJSONKey: ToJSON,
	}
}
