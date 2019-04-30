package flagsh

import (
	"cmd/asm/internal/flags"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ParseKey = "Parse"

	UsageKey = "Usage"
)

func New() hctx.Map {
	return hctx.Map{

		ParseKey: Parse,

		UsageKey: Usage,
	}
}

var Parse = flags.Parse

var Usage = flags.Usage
