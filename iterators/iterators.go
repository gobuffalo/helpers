package iterators

import "github.com/gobuffalo/helpers/hctx"

func New() hctx.Map {
	return hctx.Map{
		"range":   Range,
		"between": Between,
		"until":   Until,
		"groupBy": GroupBy,
	}
}
