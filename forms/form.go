package forms

import (
	"html/template"

	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/tags"
	"github.com/gobuffalo/tags/form"
)

// FormHelper implements a Plush helper around the
// form.New function in the github.com/gobuffalo/tags/form package
func Form(opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	return helper(opts, help, func(opts tags.Options) helperable {
		return form.New(opts)
	})
}
