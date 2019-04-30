package cacheh

import (
	"cmd/go/internal/cache"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DefaultKey = "Default"

	DefaultDirKey = "DefaultDir"

	FileHashKey = "FileHash"

	NewHashKey = "NewHash"

	OpenKey = "Open"

	SetFileHashKey = "SetFileHash"

	SubkeyKey = "Subkey"
)

func New() hctx.Map {
	return hctx.Map{

		DefaultKey: Default,

		DefaultDirKey: DefaultDir,

		FileHashKey: FileHash,

		NewHashKey: NewHash,

		OpenKey: Open,

		SetFileHashKey: SetFileHash,

		SubkeyKey: Subkey,
	}
}

// Default returns the default cache to use, or nil if no cache should be used.
var Default = cache.Default

// DefaultDir returns the effective GOCACHE setting.
// It returns &#34;off&#34; if the cache is disabled.
var DefaultDir = cache.DefaultDir

// FileHash returns the hash of the named file.
// It caches repeated lookups for a given file,
// and the cache entry for a file can be initialized
// using SetFileHash.
// The hash used by FileHash is not the same as
// the hash used by NewHash.
var FileHash = cache.FileHash

// NewHash returns a new Hash.
// The caller is expected to Write data to it and then call Sum.
var NewHash = cache.NewHash

// Open opens and returns the cache in the given directory.
//
// It is safe for multiple processes on a single machine to use the
// same cache directory in a local file system simultaneously.
// They will coordinate using operating system file locks and may
// duplicate effort but will not corrupt the cache.
//
// However, it is NOT safe for multiple processes on different machines
// to share a cache directory (for example, if the directory were stored
// in a network file system). File locking is notoriously unreliable in
// network file systems and may not suffice to protect the cache.
var Open = cache.Open

// SetFileHash sets the hash returned by FileHash for file.
var SetFileHash = cache.SetFileHash

// Subkey returns an action ID corresponding to mixing a parent
// action ID with a string description of the subkey.
var Subkey = cache.Subkey
