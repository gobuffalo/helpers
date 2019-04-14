package forms

import (
	"html/template"

	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/tags"
	"github.com/gobuffalo/tags/form"
)

// FormForimplements a Plush helper around the
// form.NewFormFor function in the github.com/gobuffalo/tags/form package
func FormFor(model interface{}, opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	return helper(opts, help, func(opts tags.Options) helperable {
		return form.NewFormFor(model, opts)
	})
}
