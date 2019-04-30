package crc32h

import (
	"hash/crc32"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ChecksumKey = "Checksum"

	ChecksumIEEEKey = "ChecksumIEEE"

	MakeTableKey = "MakeTable"

	NewKey = "New"

	NewIEEEKey = "NewIEEE"

	UpdateKey = "Update"
)

func New() hctx.Map {
	return hctx.Map{

		ChecksumKey: Checksum,

		ChecksumIEEEKey: ChecksumIEEE,

		MakeTableKey: MakeTable,

		NewKey: New,

		NewIEEEKey: NewIEEE,

		UpdateKey: Update,
	}
}

// Checksum returns the CRC-32 checksum of data
// using the polynomial represented by the Table.
var Checksum = crc32.Checksum

// ChecksumIEEE returns the CRC-32 checksum of data
// using the IEEE polynomial.
var ChecksumIEEE = crc32.ChecksumIEEE

// MakeTable returns a Table constructed from the specified polynomial.
// The contents of this Table must not be modified.
var MakeTable = crc32.MakeTable

// New creates a new hash.Hash32 computing the CRC-32 checksum using the
// polynomial represented by the Table. Its Sum method will lay the
// value out in big-endian byte order. The returned Hash32 also
// implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to
// marshal and unmarshal the internal state of the hash.
var New = crc32.New

// NewIEEE creates a new hash.Hash32 computing the CRC-32 checksum using
// the IEEE polynomial. Its Sum method will lay the value out in
// big-endian byte order. The returned Hash32 also implements
// encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to marshal
// and unmarshal the internal state of the hash.
var NewIEEE = crc32.NewIEEE

// Update returns the result of adding the bytes in p to the crc.
var Update = crc32.Update
