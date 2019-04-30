package gobh

import (
	"encoding/gob"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DebugKey = "Debug"

	NewDecoderKey = "NewDecoder"

	NewEncoderKey = "NewEncoder"

	RegisterKey = "Register"

	RegisterNameKey = "RegisterName"
)

func New() hctx.Map {
	return hctx.Map{

		DebugKey: Debug,

		NewDecoderKey: NewDecoder,

		NewEncoderKey: NewEncoder,

		RegisterKey: Register,

		RegisterNameKey: RegisterName,
	}
}

// Debug prints a human-readable representation of the gob data read from r.
// It is a no-op unless debugging was enabled when the package was built.
var Debug = gob.Debug

// NewDecoder returns a new decoder that reads from the io.Reader.
// If r does not also implement io.ByteReader, it will be wrapped in a
// bufio.Reader.
var NewDecoder = gob.NewDecoder

// NewEncoder returns a new encoder that will transmit on the io.Writer.
var NewEncoder = gob.NewEncoder

// Register records a type, identified by a value for that type, under its
// internal type name. That name will identify the concrete type of a value
// sent or received as an interface variable. Only types that will be
// transferred as implementations of interface values need to be registered.
// Expecting to be used only during initialization, it panics if the mapping
// between types and names is not a bijection.
var Register = gob.Register

// RegisterName is like Register but uses the provided name rather than the
// type&#39;s default.
var RegisterName = gob.RegisterName
