package text

import (
	"strings"

	"github.com/gobuffalo/helpers/hctx"
)

// Upcase returns an uppercased version of the given string
func Upcase(s string, opts hctx.Map) string {
	return strings.ToUpper(s)
}
