package iotesth

import (
	"testing/iotest"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DataErrReaderKey = "DataErrReader"

	HalfReaderKey = "HalfReader"

	NewReadLoggerKey = "NewReadLogger"

	NewWriteLoggerKey = "NewWriteLogger"

	OneByteReaderKey = "OneByteReader"

	TimeoutReaderKey = "TimeoutReader"

	TruncateWriterKey = "TruncateWriter"
)

func New() hctx.Map {
	return hctx.Map{

		DataErrReaderKey: DataErrReader,

		HalfReaderKey: HalfReader,

		NewReadLoggerKey: NewReadLogger,

		NewWriteLoggerKey: NewWriteLogger,

		OneByteReaderKey: OneByteReader,

		TimeoutReaderKey: TimeoutReader,

		TruncateWriterKey: TruncateWriter,
	}
}

// DataErrReader changes the way errors are handled by a Reader. Normally, a
// Reader returns an error (typically EOF) from the first Read call after the
// last piece of data is read. DataErrReader wraps a Reader and changes its
// behavior so the final error is returned along with the final data, instead
// of in the first call after the final data.
var DataErrReader = iotest.DataErrReader

// HalfReader returns a Reader that implements Read
// by reading half as many requested bytes from r.
var HalfReader = iotest.HalfReader

// NewReadLogger returns a reader that behaves like r except
// that it logs (using log.Printf) each read to standard error,
// printing the prefix and the hexadecimal data read.
var NewReadLogger = iotest.NewReadLogger

// NewWriteLogger returns a writer that behaves like w except
// that it logs (using log.Printf) each write to standard error,
// printing the prefix and the hexadecimal data written.
var NewWriteLogger = iotest.NewWriteLogger

// OneByteReader returns a Reader that implements
// each non-empty Read by reading one byte from r.
var OneByteReader = iotest.OneByteReader

// TimeoutReader returns ErrTimeout on the second read
// with no data. Subsequent calls to read succeed.
var TimeoutReader = iotest.TimeoutReader

// TruncateWriter returns a Writer that writes to w
// but stops silently after n bytes.
var TruncateWriter = iotest.TruncateWriter
