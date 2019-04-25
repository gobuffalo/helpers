package text

import (
	"testing"

	"github.com/gobuffalo/helpers/hctx"
)

func TestUpcase(t *testing.T) {
	type args struct {
		s    string
		opts hctx.Map
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Normal", args{"thiS Normal", nil}, "THIS NORMAL"},
		{"Normal 2", args{"thiS Is Normal", nil}, "THIS IS NORMAL"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Upcase(tt.args.s, tt.args.opts); got != tt.want {
				t.Errorf("Upcase() = %v, want %v", got, tt.want)
			}
		})
	}
}
