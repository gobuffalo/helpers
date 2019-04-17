package tags

import "github.com/gobuffalo/helpers/hctx"

const (
	ImgKey          = "imgTag"
	CSSKey          = "stylesheetTag"
	JSKey           = "javascriptTag"
	LinkToKey       = "linkTo"
	RemoteLinkToKey = "remoteLinkTo"
)

func New() hctx.Map {
	return hctx.Map{
		ImgKey:          Img,
		"img":           Img,
		"css":           CSS,
		"cssTag":        CSS,
		CSSKey:          CSS,
		"js":            JS,
		"jsTag":         JS,
		JSKey:           JS,
		LinkToKey:       LinkTo,
		RemoteLinkToKey: RemoteLinkTo,
	}
}
