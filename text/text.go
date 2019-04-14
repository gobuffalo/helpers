package text

import "github.com/gobuffalo/helpers/hctx"

func New() hctx.Map {
	return hctx.Map{
		"markdown": Markdown,
		"truncate": Truncate,
	}
}
