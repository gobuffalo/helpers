package bootstrap

import (
	"html/template"

	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/tags"
	"github.com/gobuffalo/tags/form/bootstrap"
)

// BootstrapFormForHelper implements a Plush helper around the
// bootstrap.NewFormFor function in the github.com/gobuffalo/tags/form/bootstrap package
func BootstrapFormFor(model interface{}, opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	return helper(opts, help, func(opts tags.Options) helperable {
		return bootstrap.NewFormFor(model, opts)
	})
}
