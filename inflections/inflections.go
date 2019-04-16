package inflections

import (
	"strings"

	"github.com/gobuffalo/flect"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	CamelizeKey    = "camelize"
	CapitalizeKey  = "capitalize"
	DasherizeKey   = "dasherize"
	OrdinalizeKey  = "ordinalize"
	PluralizeKey   = "pluralize"
	SingularizeKey = "singularize"
	UnderscoreKey  = "underscore"
)

func New() hctx.Map {
	return hctx.Map{
		CamelizeKey:           Camelize,
		"camelize_down_first": Camelize, // Deprecated
		CapitalizeKey:         Capitalize,
		DasherizeKey:          Dasherize,
		OrdinalizeKey:         Ordinalize,
		PluralizeKey:          Pluralize,
		SingularizeKey:        Singularize,
		UnderscoreKey:         Underscore,
		// "asciffy":             Asciify,
		// "humanize":            Humanize,
		// "parameterize":        Parameterize,
		// "pluralize_with_size": PluralizeWithSize,
		// "tableize":            Tableize,
		// "typeify":             Typeify,
	}
}

var Upcase = strings.ToUpper
var Downcase = strings.ToLower
var Camelize = flect.Camelize
var Pascalize = flect.Pascalize
var Capitalize = flect.Capitalize
var Dasherize = flect.Dasherize
var Ordinalize = flect.Ordinalize
var Pluralize = flect.Pluralize
var Singularize = flect.Singularize
var Underscore = flect.Underscore
