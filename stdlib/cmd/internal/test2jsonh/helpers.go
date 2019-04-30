package test2jsonh

import (
	"cmd/internal/test2json"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewConverterKey = "NewConverter"
)

func New() hctx.Map {
	return hctx.Map{

		NewConverterKey: NewConverter,
	}
}

// NewConverter returns a &#34;test to json&#34; converter.
// Writes on the returned writer are written as JSON to w,
// with minimal delay.
//
// The writes to w are whole JSON events ending in \n,
// so that it is safe to run multiple tests writing to multiple converters
// writing to a single underlying output stream w.
// As long as the underlying output w can handle concurrent writes
// from multiple goroutines, the result will be a JSON stream
// describing the relative ordering of execution in all the concurrent tests.
//
// The mode flag adjusts the behavior of the converter.
// Passing ModeTime includes event timestamps and elapsed times.
//
// The pkg string, if present, specifies the import path to
// report in the JSON stream.
var NewConverter = test2json.NewConverter
