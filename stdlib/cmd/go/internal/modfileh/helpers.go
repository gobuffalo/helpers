package modfileh

import (
	"cmd/go/internal/modfile"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AutoQuoteKey = "AutoQuote"

	FormatKey = "Format"

	IsDirectoryPathKey = "IsDirectoryPath"

	ModulePathKey = "ModulePath"

	MustQuoteKey = "MustQuote"

	ParseKey = "Parse"

	ParseGopkgInKey = "ParseGopkgIn"

	ParseLaxKey = "ParseLax"
)

func New() hctx.Map {
	return hctx.Map{

		AutoQuoteKey: AutoQuote,

		FormatKey: Format,

		IsDirectoryPathKey: IsDirectoryPath,

		ModulePathKey: ModulePath,

		MustQuoteKey: MustQuote,

		ParseKey: Parse,

		ParseGopkgInKey: ParseGopkgIn,

		ParseLaxKey: ParseLax,
	}
}

// AutoQuote returns s or, if quoting is required for s to appear in a go.mod,
// the quotation of s.
var AutoQuote = modfile.AutoQuote

var Format = modfile.Format

// IsDirectoryPath reports whether the given path should be interpreted
// as a directory path. Just like on the go command line, relative paths
// and rooted paths are directory paths; the rest are module paths.
var IsDirectoryPath = modfile.IsDirectoryPath

// ModulePath returns the module path from the gomod file text.
// If it cannot find a module path, it returns an empty string.
// It is tolerant of unrelated problems in the go.mod file.
var ModulePath = modfile.ModulePath

// MustQuote reports whether s must be quoted in order to appear as
// a single token in a go.mod line.
var MustQuote = modfile.MustQuote

// Parse parses the data, reported in errors as being from file,
// into a File struct. It applies fix, if non-nil, to canonicalize all module versions found.
var Parse = modfile.Parse

// ParseGopkgIn splits gopkg.in import paths into their constituent parts
var ParseGopkgIn = modfile.ParseGopkgIn

// ParseLax is like Parse but ignores unknown statements.
// It is used when parsing go.mod files other than the main module,
// under the theory that most statement types we add in the future will
// only apply in the main module, like exclude and replace,
// and so we get better gradual deployments if old go commands
// simply ignore those statements when found in go.mod files
// in dependencies.
var ParseLax = modfile.ParseLax
