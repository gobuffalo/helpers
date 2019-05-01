package rpch

import (
	"net/rpc"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AcceptKey = "Accept"

	DialKey = "Dial"

	DialHTTPKey = "DialHTTP"

	DialHTTPPathKey = "DialHTTPPath"

	HandleHTTPKey = "HandleHTTP"

	NewClientKey = "NewClient"

	NewClientWithCodecKey = "NewClientWithCodec"

	NewServerKey = "NewServer"

	RegisterKey = "Register"

	RegisterNameKey = "RegisterName"

	ServeCodecKey = "ServeCodec"

	ServeConnKey = "ServeConn"

	ServeRequestKey = "ServeRequest"
)

func New() hctx.Map {
	return hctx.Map{

		AcceptKey: Accept,

		DialKey: Dial,

		DialHTTPKey: DialHTTP,

		DialHTTPPathKey: DialHTTPPath,

		HandleHTTPKey: HandleHTTP,

		NewClientKey: NewClient,

		NewClientWithCodecKey: NewClientWithCodec,

		NewServerKey: NewServer,

		RegisterKey: Register,

		RegisterNameKey: RegisterName,

		ServeCodecKey: ServeCodec,

		ServeConnKey: ServeConn,

		ServeRequestKey: ServeRequest,
	}
}

// Accept accepts connections on the listener and serves requests
// to DefaultServer for each incoming connection.
// Accept blocks; the caller typically invokes it in a go statement.
var Accept = rpc.Accept

// Dial connects to an RPC server at the specified network address.
var Dial = rpc.Dial

// DialHTTP connects to an HTTP RPC server at the specified network address
// listening on the default HTTP RPC path.
var DialHTTP = rpc.DialHTTP

// DialHTTPPath connects to an HTTP RPC server
// at the specified network address and path.
var DialHTTPPath = rpc.DialHTTPPath

// HandleHTTP registers an HTTP handler for RPC messages to DefaultServer
// on DefaultRPCPath and a debugging handler on DefaultDebugPath.
// It is still necessary to invoke http.Serve(), typically in a go statement.
var HandleHTTP = rpc.HandleHTTP

// NewClient returns a new Client to handle requests to the
// set of services at the other end of the connection.
// It adds a buffer to the write side of the connection so
// the header and payload are sent as a unit.
var NewClient = rpc.NewClient

// NewClientWithCodec is like NewClient but uses the specified
// codec to encode requests and decode responses.
var NewClientWithCodec = rpc.NewClientWithCodec

// NewServer returns a new Server.
var NewServer = rpc.NewServer

// Register publishes the receiver&#39;s methods in the DefaultServer.
var Register = rpc.Register

// RegisterName is like Register but uses the provided name for the type
// instead of the receiver&#39;s concrete type.
var RegisterName = rpc.RegisterName

// ServeCodec is like ServeConn but uses the specified codec to
// decode requests and encode responses.
var ServeCodec = rpc.ServeCodec

// ServeConn runs the DefaultServer on a single connection.
// ServeConn blocks, serving the connection until the client hangs up.
// The caller typically invokes ServeConn in a go statement.
// ServeConn uses the gob wire format (see package gob) on the
// connection. To use an alternate codec, use ServeCodec.
var ServeConn = rpc.ServeConn

// ServeRequest is like ServeCodec but synchronously serves a single request.
// It does not close the codec upon completion.
var ServeRequest = rpc.ServeRequest
