package ioh

import (
	"io"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	CopyKey = "Copy"

	CopyBufferKey = "CopyBuffer"

	CopyNKey = "CopyN"

	LimitReaderKey = "LimitReader"

	MultiReaderKey = "MultiReader"

	MultiWriterKey = "MultiWriter"

	NewSectionReaderKey = "NewSectionReader"

	PipeKey = "Pipe"

	ReadAtLeastKey = "ReadAtLeast"

	ReadFullKey = "ReadFull"

	TeeReaderKey = "TeeReader"

	WriteStringKey = "WriteString"
)

func New() hctx.Map {
	return hctx.Map{

		CopyKey: Copy,

		CopyBufferKey: CopyBuffer,

		CopyNKey: CopyN,

		LimitReaderKey: LimitReader,

		MultiReaderKey: MultiReader,

		MultiWriterKey: MultiWriter,

		NewSectionReaderKey: NewSectionReader,

		PipeKey: Pipe,

		ReadAtLeastKey: ReadAtLeast,

		ReadFullKey: ReadFull,

		TeeReaderKey: TeeReader,

		WriteStringKey: WriteString,
	}
}

// Copy copies from src to dst until either EOF is reached
// on src or an error occurs. It returns the number of bytes
// copied and the first error encountered while copying, if any.
//
// A successful Copy returns err == nil, not err == EOF.
// Because Copy is defined to read from src until EOF, it does
// not treat an EOF from Read as an error to be reported.
//
// If src implements the WriterTo interface,
// the copy is implemented by calling src.WriteTo(dst).
// Otherwise, if dst implements the ReaderFrom interface,
// the copy is implemented by calling dst.ReadFrom(src).
var Copy = io.Copy

// CopyBuffer is identical to Copy except that it stages through the
// provided buffer (if one is required) rather than allocating a
// temporary one. If buf is nil, one is allocated; otherwise if it has
// zero length, CopyBuffer panics.
var CopyBuffer = io.CopyBuffer

// CopyN copies n bytes (or until an error) from src to dst.
// It returns the number of bytes copied and the earliest
// error encountered while copying.
// On return, written == n if and only if err == nil.
//
// If dst implements the ReaderFrom interface,
// the copy is implemented using it.
var CopyN = io.CopyN

// LimitReader returns a Reader that reads from r
// but stops with EOF after n bytes.
// The underlying implementation is a *LimitedReader.
var LimitReader = io.LimitReader

// MultiReader returns a Reader that&#39;s the logical concatenation of
// the provided input readers. They&#39;re read sequentially. Once all
// inputs have returned EOF, Read will return EOF.  If any of the readers
// return a non-nil, non-EOF error, Read will return that error.
var MultiReader = io.MultiReader

// MultiWriter creates a writer that duplicates its writes to all the
// provided writers, similar to the Unix tee(1) command.
//
// Each write is written to each listed writer, one at a time.
// If a listed writer returns an error, that overall write operation
// stops and returns the error; it does not continue down the list.
var MultiWriter = io.MultiWriter

// NewSectionReader returns a SectionReader that reads from r
// starting at offset off and stops with EOF after n bytes.
var NewSectionReader = io.NewSectionReader

// Pipe creates a synchronous in-memory pipe.
// It can be used to connect code expecting an io.Reader
// with code expecting an io.Writer.
//
// Reads and Writes on the pipe are matched one to one
// except when multiple Reads are needed to consume a single Write.
// That is, each Write to the PipeWriter blocks until it has satisfied
// one or more Reads from the PipeReader that fully consume
// the written data.
// The data is copied directly from the Write to the corresponding
// Read (or Reads); there is no internal buffering.
//
// It is safe to call Read and Write in parallel with each other or with Close.
// Parallel calls to Read and parallel calls to Write are also safe:
// the individual calls will be gated sequentially.
var Pipe = io.Pipe

// ReadAtLeast reads from r into buf until it has read at least min bytes.
// It returns the number of bytes copied and an error if fewer bytes were read.
// The error is EOF only if no bytes were read.
// If an EOF happens after reading fewer than min bytes,
// ReadAtLeast returns ErrUnexpectedEOF.
// If min is greater than the length of buf, ReadAtLeast returns ErrShortBuffer.
// On return, n &gt;= min if and only if err == nil.
var ReadAtLeast = io.ReadAtLeast

// ReadFull reads exactly len(buf) bytes from r into buf.
// It returns the number of bytes copied and an error if fewer bytes were read.
// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// ReadFull returns ErrUnexpectedEOF.
// On return, n == len(buf) if and only if err == nil.
var ReadFull = io.ReadFull

// TeeReader returns a Reader that writes to w what it reads from r.
// All reads from r performed through it are matched with
// corresponding writes to w. There is no internal buffering -
// the write must complete before the read completes.
// Any error encountered while writing is reported as a read error.
var TeeReader = io.TeeReader

// WriteString writes the contents of the string s to w, which accepts a slice of bytes.
// If w implements a WriteString method, it is invoked directly.
// Otherwise, w.Write is called exactly once.
var WriteString = io.WriteString
