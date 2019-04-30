package hpackh

import (
	"internal/x/net/http2/hpack"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AppendHuffmanStringKey = "AppendHuffmanString"

	HuffmanDecodeKey = "HuffmanDecode"

	HuffmanDecodeToStringKey = "HuffmanDecodeToString"

	HuffmanEncodeLengthKey = "HuffmanEncodeLength"

	NewDecoderKey = "NewDecoder"

	NewEncoderKey = "NewEncoder"
)

func New() hctx.Map {
	return hctx.Map{

		AppendHuffmanStringKey: AppendHuffmanString,

		HuffmanDecodeKey: HuffmanDecode,

		HuffmanDecodeToStringKey: HuffmanDecodeToString,

		HuffmanEncodeLengthKey: HuffmanEncodeLength,

		NewDecoderKey: NewDecoder,

		NewEncoderKey: NewEncoder,
	}
}

// AppendHuffmanString appends s, as encoded in Huffman codes, to dst
// and returns the extended buffer.
var AppendHuffmanString = hpack.AppendHuffmanString

// HuffmanDecode decodes the string in v and writes the expanded
// result to w, returning the number of bytes written to w and the
// Write call&#39;s return value. At most one Write call is made.
var HuffmanDecode = hpack.HuffmanDecode

// HuffmanDecodeToString decodes the string in v.
var HuffmanDecodeToString = hpack.HuffmanDecodeToString

// HuffmanEncodeLength returns the number of bytes required to encode
// s in Huffman codes. The result is round up to byte boundary.
var HuffmanEncodeLength = hpack.HuffmanEncodeLength

// NewDecoder returns a new decoder with the provided maximum dynamic
// table size. The emitFunc will be called for each valid field
// parsed, in the same goroutine as calls to Write, before Write returns.
var NewDecoder = hpack.NewDecoder

// NewEncoder returns a new Encoder which performs HPACK encoding. An
// encoded data is written to w.
var NewEncoder = hpack.NewEncoder
