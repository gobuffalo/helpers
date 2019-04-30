package browserh

import (
	"cmd/internal/browser"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	CommandsKey = "Commands"

	OpenKey = "Open"
)

func New() hctx.Map {
	return hctx.Map{

		CommandsKey: Commands,

		OpenKey: Open,
	}
}

// Commands returns a list of possible commands to use to open a url.
var Commands = browser.Commands

// Open tries to open url in a browser and reports whether it succeeded.
var Open = browser.Open
