package crc64h

import (
	"hash/crc64"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ChecksumKey = "Checksum"

	MakeTableKey = "MakeTable"

	UpdateKey = "Update"
)

func New() hctx.Map {
	return hctx.Map{

		ChecksumKey: Checksum,

		MakeTableKey: MakeTable,

		UpdateKey: Update,
	}
}

// Checksum returns the CRC-64 checksum of data
// using the polynomial represented by the Table.
var Checksum = crc64.Checksum

// MakeTable returns a Table constructed from the specified polynomial.
// The contents of this Table must not be modified.
var MakeTable = crc64.MakeTable

// Update returns the result of adding the bytes in p to the crc.
var Update = crc64.Update
