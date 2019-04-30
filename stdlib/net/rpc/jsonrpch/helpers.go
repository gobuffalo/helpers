package jsonrpch

import (
	"net/rpc/jsonrpc"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DialKey = "Dial"

	NewClientKey = "NewClient"

	NewClientCodecKey = "NewClientCodec"

	NewServerCodecKey = "NewServerCodec"

	ServeConnKey = "ServeConn"
)

func New() hctx.Map {
	return hctx.Map{

		DialKey: Dial,

		NewClientKey: NewClient,

		NewClientCodecKey: NewClientCodec,

		NewServerCodecKey: NewServerCodec,

		ServeConnKey: ServeConn,
	}
}

// Dial connects to a JSON-RPC server at the specified network address.
var Dial = jsonrpc.Dial

// NewClient returns a new rpc.Client to handle requests to the
// set of services at the other end of the connection.
var NewClient = jsonrpc.NewClient

// NewClientCodec returns a new rpc.ClientCodec using JSON-RPC on conn.
var NewClientCodec = jsonrpc.NewClientCodec

// NewServerCodec returns a new rpc.ServerCodec using JSON-RPC on conn.
var NewServerCodec = jsonrpc.NewServerCodec

// ServeConn runs the JSON-RPC server on a single connection.
// ServeConn blocks, serving the connection until the client hangs up.
// The caller typically invokes ServeConn in a go statement.
var ServeConn = jsonrpc.ServeConn
