package webtesth

import (
	"cmd/go/internal/webtest"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DoKey = "Do"

	DoPrintKey = "DoPrint"

	HookKey = "Hook"

	LoadOnceKey = "LoadOnce"

	PrintKey = "Print"

	ServeKey = "Serve"

	UnhookKey = "Unhook"
)

func New() hctx.Map {
	return hctx.Map{

		DoKey: Do,

		DoPrintKey: DoPrint,

		HookKey: Hook,

		LoadOnceKey: LoadOnce,

		PrintKey: Print,

		ServeKey: Serve,

		UnhookKey: Unhook,
	}
}

var Do = webtest.Do

var DoPrint = webtest.DoPrint

var Hook = webtest.Hook

var LoadOnce = webtest.LoadOnce

var Print = webtest.Print

var Serve = webtest.Serve

var Unhook = webtest.Unhook
