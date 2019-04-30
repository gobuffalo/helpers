package buildidh

import (
	"cmd/internal/buildid"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	FindAndHashKey = "FindAndHash"

	ReadELFNoteKey = "ReadELFNote"

	ReadFileKey = "ReadFile"

	RewriteKey = "Rewrite"
)

func New() hctx.Map {
	return hctx.Map{

		FindAndHashKey: FindAndHash,

		ReadELFNoteKey: ReadELFNote,

		ReadFileKey: ReadFile,

		RewriteKey: Rewrite,
	}
}

// FindAndHash reads all of r and returns the offsets of occurrences of id.
// While reading, findAndHash also computes and returns
// a hash of the content of r, but with occurrences of id replaced by zeros.
// FindAndHash reads bufSize bytes from r at a time.
// If bufSize == 0, FindAndHash uses a reasonable default.
var FindAndHash = buildid.FindAndHash

var ReadELFNote = buildid.ReadELFNote

// ReadFile reads the build ID from an archive or executable file.
var ReadFile = buildid.ReadFile

var Rewrite = buildid.Rewrite
