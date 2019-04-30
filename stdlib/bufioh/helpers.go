package bufioh

import (
	"bufio"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewReadWriterKey = "NewReadWriter"

	NewReaderKey = "NewReader"

	NewReaderSizeKey = "NewReaderSize"

	NewScannerKey = "NewScanner"

	NewWriterKey = "NewWriter"

	NewWriterSizeKey = "NewWriterSize"

	ScanBytesKey = "ScanBytes"

	ScanLinesKey = "ScanLines"

	ScanRunesKey = "ScanRunes"

	ScanWordsKey = "ScanWords"
)

func New() hctx.Map {
	return hctx.Map{

		NewReadWriterKey: NewReadWriter,

		NewReaderKey: NewReader,

		NewReaderSizeKey: NewReaderSize,

		NewScannerKey: NewScanner,

		NewWriterKey: NewWriter,

		NewWriterSizeKey: NewWriterSize,

		ScanBytesKey: ScanBytes,

		ScanLinesKey: ScanLines,

		ScanRunesKey: ScanRunes,

		ScanWordsKey: ScanWords,
	}
}

// NewReadWriter allocates a new ReadWriter that dispatches to r and w.
var NewReadWriter = bufio.NewReadWriter

// NewReader returns a new Reader whose buffer has the default size.
var NewReader = bufio.NewReader

// NewReaderSize returns a new Reader whose buffer has at least the specified
// size. If the argument io.Reader is already a Reader with large enough
// size, it returns the underlying Reader.
var NewReaderSize = bufio.NewReaderSize

// NewScanner returns a new Scanner to read from r.
// The split function defaults to ScanLines.
var NewScanner = bufio.NewScanner

// NewWriter returns a new Writer whose buffer has the default size.
var NewWriter = bufio.NewWriter

// NewWriterSize returns a new Writer whose buffer has at least the specified
// size. If the argument io.Writer is already a Writer with large enough
// size, it returns the underlying Writer.
var NewWriterSize = bufio.NewWriterSize

// ScanBytes is a split function for a Scanner that returns each byte as a token.
var ScanBytes = bufio.ScanBytes

// ScanLines is a split function for a Scanner that returns each line of
// text, stripped of any trailing end-of-line marker. The returned line may
// be empty. The end-of-line marker is one optional carriage return followed
// by one mandatory newline. In regular expression notation, it is `\r?\n`.
// The last non-empty line of input will be returned even if it has no
// newline.
var ScanLines = bufio.ScanLines

// ScanRunes is a split function for a Scanner that returns each
// UTF-8-encoded rune as a token. The sequence of runes returned is
// equivalent to that from a range loop over the input as a string, which
// means that erroneous UTF-8 encodings translate to U+FFFD = &#34;\xef\xbf\xbd&#34;.
// Because of the Scan interface, this makes it impossible for the client to
// distinguish correctly encoded replacement runes from encoding errors.
var ScanRunes = bufio.ScanRunes

// ScanWords is a split function for a Scanner that returns each
// space-separated word of text, with surrounding spaces deleted. It will
// never return an empty string. The definition of space is set by
// unicode.IsSpace.
var ScanWords = bufio.ScanWords
