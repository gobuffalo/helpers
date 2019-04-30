package gcimporterh

import (
	"go/internal/gcimporter"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	BImportDataKey = "BImportData"

	FindExportDataKey = "FindExportData"

	FindPkgKey = "FindPkg"

	ImportKey = "Import"
)

func New() hctx.Map {
	return hctx.Map{

		BImportDataKey: BImportData,

		FindExportDataKey: FindExportData,

		FindPkgKey: FindPkg,

		ImportKey: Import,
	}
}

// BImportData imports a package from the serialized package data
// and returns the number of bytes consumed and a reference to the package.
// If the export data version is not recognized or the format is otherwise
// compromised, an error is returned.
var BImportData = gcimporter.BImportData

// FindExportData positions the reader r at the beginning of the
// export data section of an underlying GC-created object/archive
// file by reading from it. The reader must be positioned at the
// start of the file before calling this function. The hdr result
// is the string before the export data, either &#34;$$&#34; or &#34;$$B&#34;.
var FindExportData = gcimporter.FindExportData

// FindPkg returns the filename and unique package id for an import
// path based on package information provided by build.Import (using
// the build.Default build.Context). A relative srcDir is interpreted
// relative to the current working directory.
// If no file was found, an empty filename is returned.
var FindPkg = gcimporter.FindPkg

// Import imports a gc-generated package given its import path and srcDir, adds
// the corresponding package object to the packages map, and returns the object.
// The packages map must contain all packages already imported.
var Import = gcimporter.Import
