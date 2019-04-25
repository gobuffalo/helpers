package text

import (
	"testing"

	"github.com/gobuffalo/helpers/hctx"
)

func TestDowncase(t *testing.T) {
	type args struct {
		s    string
		opts hctx.Map
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Normal", args{"thiS Normal", nil}, "this normal"},
		{"Normal 2", args{"thiS Is Normal", nil}, "this is normal"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Downcase(tt.args.s, tt.args.opts); got != tt.want {
				t.Errorf("Downcase() = %v, want %v", got, tt.want)
			}
		})
	}
}
