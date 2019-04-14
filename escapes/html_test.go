package escapes

import (
	"html/template"
	"testing"

	"github.com/gobuffalo/helpers/helptest"
	"github.com/stretchr/testify/require"
)

func Test_HTMLEscape(t *testing.T) {
	r := require.New(t)

	in := `<a href="/path?foo=Bar">foo</a>`
	hc := helptest.NewContext()
	s, err := HTMLEscape(in, hc)
	r.NoError(err)
	r.Equal(template.HTMLEscapeString(in), s)
}

func Test_HTMLEscape_Block(t *testing.T) {
	r := require.New(t)

	in := `<a href="/path?foo=Bar">foo</a>`
	hc := helptest.NewContext()
	hc.BlockFn = func() (string, error) {
		return in, nil
	}
	s, err := HTMLEscape("", hc)
	r.NoError(err)
	r.Equal(template.HTMLEscapeString(in), s)
}
