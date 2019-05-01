package hexh

import (
	"encoding/hex"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DecodeKey = "Decode"

	DecodeStringKey = "DecodeString"

	DecodedLenKey = "DecodedLen"

	DumpKey = "Dump"

	DumperKey = "Dumper"

	EncodeKey = "Encode"

	EncodeToStringKey = "EncodeToString"

	EncodedLenKey = "EncodedLen"

	NewDecoderKey = "NewDecoder"

	NewEncoderKey = "NewEncoder"
)

func New() hctx.Map {
	return hctx.Map{

		DecodeKey: Decode,

		DecodeStringKey: DecodeString,

		DecodedLenKey: DecodedLen,

		DumpKey: Dump,

		DumperKey: Dumper,

		EncodeKey: Encode,

		EncodeToStringKey: EncodeToString,

		EncodedLenKey: EncodedLen,

		NewDecoderKey: NewDecoder,

		NewEncoderKey: NewEncoder,
	}
}

// Decode decodes src into DecodedLen(len(src)) bytes,
// returning the actual number of bytes written to dst.
//
// Decode expects that src contain only hexadecimal
// characters and that src should have an even length.
// If the input is malformed, Decode returns the number
// of bytes decoded before the error.
var Decode = hex.Decode

// DecodeString returns the bytes represented by the hexadecimal string s.
//
// DecodeString expects that src contain only hexadecimal
// characters and that src should have an even length.
// If the input is malformed, DecodeString returns a string
// containing the bytes decoded before the error.
var DecodeString = hex.DecodeString

// DecodedLen returns the length of a decoding of x source bytes.
// Specifically, it returns x / 2.
var DecodedLen = hex.DecodedLen

// Dump returns a string that contains a hex dump of the given data. The format
// of the hex dump matches the output of `hexdump -C` on the command line.
var Dump = hex.Dump

// Dumper returns a WriteCloser that writes a hex dump of all written data to
// w. The format of the dump matches the output of `hexdump -C` on the command
// line.
var Dumper = hex.Dumper

// Encode encodes src into EncodedLen(len(src))
// bytes of dst. As a convenience, it returns the number
// of bytes written to dst, but this value is always EncodedLen(len(src)).
// Encode implements hexadecimal encoding.
var Encode = hex.Encode

// EncodeToString returns the hexadecimal encoding of src.
var EncodeToString = hex.EncodeToString

// EncodedLen returns the length of an encoding of n source bytes.
// Specifically, it returns n * 2.
var EncodedLen = hex.EncodedLen

// NewDecoder returns an io.Reader that decodes hexadecimal characters from r.
// NewDecoder expects that r contain only an even number of hexadecimal characters.
var NewDecoder = hex.NewDecoder

// NewEncoder returns an io.Writer that writes lowercase hexadecimal characters to w.
var NewEncoder = hex.NewEncoder
