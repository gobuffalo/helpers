package tags

import "github.com/gobuffalo/helpers/hctx"

const (
	ImgKey = "img"
	CSSKey = "stylesheetTag"
	JSKey  = "javascriptTag"
)

func New() hctx.Map {
	return hctx.Map{
		ImgKey:   Img,
		"imgTag": Img,
		"css":    CSS,
		"cssTag": CSS,
		CSSKey:   CSS,
		"js":     JS,
		"jsTag":  JS,
		JSKey:    JS,
	}
}
