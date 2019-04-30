package txtarh

import (
	"cmd/go/internal/txtar"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	FormatKey = "Format"

	ParseKey = "Parse"

	ParseFileKey = "ParseFile"
)

func New() hctx.Map {
	return hctx.Map{

		FormatKey: Format,

		ParseKey: Parse,

		ParseFileKey: ParseFile,
	}
}

// Format returns the serialized form of an Archive.
// It is assumed that the Archive data structure is well-formed:
// a.Comment and all a.File[i].Data contain no file marker lines,
// and all a.File[i].Name is non-empty.
var Format = txtar.Format

// Parse parses the serialized form of an Archive.
// The returned Archive holds slices of data.
var Parse = txtar.Parse

// ParseFile parses the named file as an archive.
var ParseFile = txtar.ParseFile
