package expvarh

import (
	"expvar"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DoKey = "Do"

	GetKey = "Get"

	HandlerKey = "Handler"

	NewFloatKey = "NewFloat"

	NewIntKey = "NewInt"

	NewMapKey = "NewMap"

	NewStringKey = "NewString"

	PublishKey = "Publish"
)

func New() hctx.Map {
	return hctx.Map{

		DoKey: Do,

		GetKey: Get,

		HandlerKey: Handler,

		NewFloatKey: NewFloat,

		NewIntKey: NewInt,

		NewMapKey: NewMap,

		NewStringKey: NewString,

		PublishKey: Publish,
	}
}

// Do calls f for each exported variable.
// The global variable map is locked during the iteration,
// but existing entries may be concurrently updated.
var Do = expvar.Do

// Get retrieves a named exported variable. It returns nil if the name has
// not been registered.
var Get = expvar.Get

// Handler returns the expvar HTTP Handler.
//
// This is only needed to install the handler in a non-standard location.
var Handler = expvar.Handler

var NewFloat = expvar.NewFloat

var NewInt = expvar.NewInt

var NewMap = expvar.NewMap

var NewString = expvar.NewString

// Publish declares a named exported variable. This should be called from a
// package&#39;s init function when it creates its Vars. If the name is already
// registered then this will log.Panic.
var Publish = expvar.Publish
