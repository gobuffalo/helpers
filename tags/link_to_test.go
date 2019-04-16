package tags

import (
	"html/template"
	"testing"

	"github.com/gobuffalo/helpers/helptest"
	"github.com/gobuffalo/tags"
	"github.com/stretchr/testify/require"
)

func Test_HREF(t *testing.T) {
	table := []struct {
		in   interface{}
		out  string
		opts tags.Options
		body string
		err  bool
	}{
		{"foo", `<a href="/foo"></a>`, tags.Options{}, "", false},
		{"foo", `<a class="btn" href="/foo"></a>`, tags.Options{"class": "btn"}, "", false},
		{[]string{"foo", "bar"}, `<a href="/foo/bar">baz</a>`, tags.Options{"body": "baz"}, "", false},
		{"foo", `<a href="/foo">my body</a>`, tags.Options{}, "my body", false},
	}

	for _, tt := range table {
		t.Run(tt.out, func(st *testing.T) {
			r := require.New(st)
			c := helptest.NewContext()
			if len(tt.body) != 0 {
				c.BlockFn = func() (string, error) {
					return tt.body, nil
				}
			}
			s, err := LinkTo(tt.in, tt.opts, c)
			if tt.err {
				r.Error(err)
				return
			}
			r.NoError(err)
			r.Equal(template.HTML(tt.out), s)
		})
	}
}
