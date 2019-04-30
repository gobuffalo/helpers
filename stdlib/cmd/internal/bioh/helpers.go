package bioh

import (
	"cmd/internal/bio"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	CreateKey = "Create"

	MustCloseKey = "MustClose"

	MustWriterKey = "MustWriter"

	OpenKey = "Open"
)

func New() hctx.Map {
	return hctx.Map{

		CreateKey: Create,

		MustCloseKey: MustClose,

		MustWriterKey: MustWriter,

		OpenKey: Open,
	}
}

// Create creates the file named name and returns a Writer
// for that file.
var Create = bio.Create

// MustClose closes Closer c and calls log.Fatal if it returns a non-nil error.
var MustClose = bio.MustClose

// MustWriter returns a Writer that wraps the provided Writer,
// except that it calls log.Fatal instead of returning a non-nil error.
var MustWriter = bio.MustWriter

// Open returns a Reader for the file named name.
var Open = bio.Open
