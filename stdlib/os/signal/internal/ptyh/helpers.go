package ptyh

import (
	"os/signal/internal/pty"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	OpenKey = "Open"
)

func New() hctx.Map {
	return hctx.Map{

		OpenKey: Open,
	}
}

// Open returns a master pty and the name of the linked slave tty.
var Open = pty.Open
