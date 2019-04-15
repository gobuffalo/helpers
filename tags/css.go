package tags

import (
	"html/template"

	"github.com/gobuffalo/tags"
)

func CSS(href string, options tags.Options) template.HTML {
	if options["rel"] == nil {
		options["rel"] = "stylesheet"
	}

	if options["media"] == nil {
		options["media"] = "screen"
	}

	options["href"] = href
	cssTag := tags.New("link", options)

	return cssTag.HTML()
}
