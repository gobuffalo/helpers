package content

import "github.com/gobuffalo/plush/v5/helpers/hctx"

// WithDefault returns the key if exists, otherwise it returns
// defaultValue passed.
func WithDefault(key string, defaultValue interface{}, help hctx.HelperContext) interface{} {
	value := help.Value(key)
	if value != nil {
		return value
	}

	return defaultValue
}
