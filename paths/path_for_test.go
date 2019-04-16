package paths

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Car struct {
	ID int
}

type boat struct {
	Slug string
}

type plane struct{}

func (plane) ToParam() string {
	return "aeroplane"
}

type truck struct{}

func (truck) ToPath() string {
	return "/a/truck"
}

func Test_PathFor(t *testing.T) {
	table := []struct {
		in  interface{}
		out string
		err bool
	}{
		{Car{1}, "/cars/1", false},
		{boat{"titanic"}, "/boats/titanic", false},
		{plane{}, "/planes/aeroplane", false},
		{truck{}, "/a/truck", false},
		{[]interface{}{truck{}, plane{}}, "/a/truck/planes/aeroplane", false},
		{"foo", "/foo", false},
		{map[int]int{}, "", true},
	}

	for _, tt := range table {
		t.Run(tt.out, func(st *testing.T) {
			r := require.New(st)
			s, err := PathFor(tt.in)
			if tt.err {
				r.Error(err)
				return
			}
			r.NoError(err)
			r.Equal(tt.out, s)
		})
	}
}
