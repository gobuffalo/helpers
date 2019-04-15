package tags

import "github.com/gobuffalo/helpers/hctx"

func New() hctx.Map {
	return hctx.Map{
		"img":    Img,
		"imgTag": Img,
		"css":    CSS,
		"cssTag": CSS,
		"js":     JS,
		"jsTag":  JS,
	}
}
