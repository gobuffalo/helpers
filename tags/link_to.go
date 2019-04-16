package tags

import (
	"html/template"

	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/helpers/paths"
	"github.com/gobuffalo/tags"
)

func LinkTo(in interface{}, opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	s, err := paths.PathFor(in)
	if err != nil {
		return "", err
	}
	opts["href"] = s
	a := tags.New("a", opts)
	if help.HasBlock() {
		x, err := help.Block()
		if err != nil {
			return "", err
		}
		a.Append(x)
	}

	return a.HTML(), nil
}
