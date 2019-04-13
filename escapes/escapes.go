package escapes

import "github.com/gobuffalo/helpers/hctx"

func New() hctx.Map {
	return hctx.Map{
		"jsEscape":   JSEscape,
		"htmlEscape": HTMLEscape,
	}
}
