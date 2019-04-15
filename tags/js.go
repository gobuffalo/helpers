package tags

import (
	"html/template"

	"github.com/gobuffalo/tags"
)

func JS(src string, options tags.Options) template.HTML {
	if options["type"] == nil {
		options["type"] = "text/javascript"
	}

	options["src"] = src
	jsTag := tags.New("script", options)

	return jsTag.HTML()
}
