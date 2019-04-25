package text

import (
	"strings"

	"github.com/gobuffalo/helpers/hctx"
)

// Downcase returns an lowercased version of the given string
func Downcase(s string, opts hctx.Map) string {
	return strings.ToLower(s)
}
