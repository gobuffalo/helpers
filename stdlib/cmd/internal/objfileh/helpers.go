package objfileh

import (
	"cmd/internal/objfile"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewFileCacheKey = "NewFileCache"

	OpenKey = "Open"
)

func New() hctx.Map {
	return hctx.Map{

		NewFileCacheKey: NewFileCache,

		OpenKey: Open,
	}
}

// NewFileCache returns a FileCache which can contain up to maxLen cached file contents.
var NewFileCache = objfile.NewFileCache

// Open opens the named file.
// The caller must call f.Close when the file is no longer needed.
var Open = objfile.Open
