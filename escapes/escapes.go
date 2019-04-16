package escapes

import "github.com/gobuffalo/helpers/hctx"

const (
	JSEscapeKey   = "jsEscape"
	HTMLEscapeKey = "htmlEscape"
)

func New() hctx.Map {
	return hctx.Map{
		JSEscapeKey:   JSEscape,
		HTMLEscapeKey: HTMLEscape,
	}
}
