package tags

import (
	"html/template"

	"github.com/gobuffalo/tags"
)

func Img(src string, options tags.Options) template.HTML {
	options["src"] = src
	imgTag := tags.New("img", options)

	return imgTag.HTML()
}
