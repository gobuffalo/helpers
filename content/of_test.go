package content

import (
	"html/template"
	"testing"

	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/helpers/helptest"
	"github.com/stretchr/testify/require"
)

func Test_ContentOf_MissingBlock(t *testing.T) {
	r := require.New(t)

	cf := helptest.NewContext()
	_, err := ContentOf("buttons", hctx.Map{}, cf)
	r.Error(err)
}

func Test_ContentOf_MissingBlock_DefaultBlock(t *testing.T) {
	r := require.New(t)

	cf := helptest.NewContext()
	cf.BlockContextFn = func(hctx.Context) (string, error) {
		return "default", nil
	}

	s, err := ContentOf("buttons", hctx.Map{}, cf)
	r.NoError(err)
	r.Equal(s, template.HTML("default"))
}
