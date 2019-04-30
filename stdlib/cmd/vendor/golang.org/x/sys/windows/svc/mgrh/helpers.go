package mgrh

import (
	"cmd/vendor/golang.org/x/sys/windows/svc/mgr"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	ConnectKey = "Connect"

	ConnectRemoteKey = "ConnectRemote"
)

func New() hctx.Map {
	return hctx.Map{

		ConnectKey: Connect,

		ConnectRemoteKey: ConnectRemote,
	}
}

// Connect establishes a connection to the service control manager.
var Connect = mgr.Connect

// ConnectRemote establishes a connection to the
// service control manager on computer named host.
var ConnectRemote = mgr.ConnectRemote
