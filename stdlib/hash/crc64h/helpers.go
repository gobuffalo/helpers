package crc64h

import (
	"hash/crc64"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ChecksumKey = "Checksum"

	MakeTableKey = "MakeTable"

	NewKey = "New"

	UpdateKey = "Update"
)

func New() hctx.Map {
	return hctx.Map{

		ChecksumKey: Checksum,

		MakeTableKey: MakeTable,

		NewKey: New,

		UpdateKey: Update,
	}
}

// Checksum returns the CRC-64 checksum of data
// using the polynomial represented by the Table.
var Checksum = crc64.Checksum

// MakeTable returns a Table constructed from the specified polynomial.
// The contents of this Table must not be modified.
var MakeTable = crc64.MakeTable

// New creates a new hash.Hash64 computing the CRC-64 checksum using the
// polynomial represented by the Table. Its Sum method will lay the
// value out in big-endian byte order. The returned Hash64 also
// implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to
// marshal and unmarshal the internal state of the hash.
var New = crc64.New

// Update returns the result of adding the bytes in p to the crc.
var Update = crc64.Update
