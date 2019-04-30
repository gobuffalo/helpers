package buildh

import (
	"go/build"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ArchCharKey = "ArchChar"

	ImportKey = "Import"

	ImportDirKey = "ImportDir"

	IsLocalImportKey = "IsLocalImport"
)

func New() hctx.Map {
	return hctx.Map{

		ArchCharKey: ArchChar,

		ImportKey: Import,

		ImportDirKey: ImportDir,

		IsLocalImportKey: IsLocalImport,
	}
}

// ArchChar returns &#34;?&#34; and an error.
// In earlier versions of Go, the returned string was used to derive
// the compiler and linker tool names, the default object file suffix,
// and the default linker output name. As of Go 1.5, those strings
// no longer vary by architecture; they are compile, link, .o, and a.out, respectively.
var ArchChar = build.ArchChar

// Import is shorthand for Default.Import.
var Import = build.Import

// ImportDir is shorthand for Default.ImportDir.
var ImportDir = build.ImportDir

// IsLocalImport reports whether the import path is
// a local import path, like &#34;.&#34;, &#34;..&#34;, &#34;./foo&#34;, or &#34;../foo&#34;.
var IsLocalImport = build.IsLocalImport
