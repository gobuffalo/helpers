package modfetchh

import (
	"cmd/go/internal/modfetch"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	CachePathKey = "CachePath"

	DownloadKey = "Download"

	DownloadDirKey = "DownloadDir"

	DownloadZipKey = "DownloadZip"

	GoModKey = "GoMod"

	GoModFileKey = "GoModFile"

	GoModSumKey = "GoModSum"

	ImportRepoRevKey = "ImportRepoRev"

	InfoFileKey = "InfoFile"

	IsPseudoVersionKey = "IsPseudoVersion"

	LookupKey = "Lookup"

	PseudoVersionKey = "PseudoVersion"

	PseudoVersionRevKey = "PseudoVersionRev"

	PseudoVersionTimeKey = "PseudoVersionTime"

	RemoveAllKey = "RemoveAll"

	SideLockKey = "SideLock"

	SortVersionsKey = "SortVersions"

	StatKey = "Stat"

	SumKey = "Sum"

	TrimGoSumKey = "TrimGoSum"

	UnzipKey = "Unzip"

	WriteGoSumKey = "WriteGoSum"
)

func New() hctx.Map {
	return hctx.Map{

		CachePathKey: CachePath,

		DownloadKey: Download,

		DownloadDirKey: DownloadDir,

		DownloadZipKey: DownloadZip,

		GoModKey: GoMod,

		GoModFileKey: GoModFile,

		GoModSumKey: GoModSum,

		ImportRepoRevKey: ImportRepoRev,

		InfoFileKey: InfoFile,

		IsPseudoVersionKey: IsPseudoVersion,

		LookupKey: Lookup,

		PseudoVersionKey: PseudoVersion,

		PseudoVersionRevKey: PseudoVersionRev,

		PseudoVersionTimeKey: PseudoVersionTime,

		RemoveAllKey: RemoveAll,

		SideLockKey: SideLock,

		SortVersionsKey: SortVersions,

		StatKey: Stat,

		SumKey: Sum,

		TrimGoSumKey: TrimGoSum,

		UnzipKey: Unzip,

		WriteGoSumKey: WriteGoSum,
	}
}

var CachePath = modfetch.CachePath

// Download downloads the specific module version to the
// local download cache and returns the name of the directory
// corresponding to the root of the module&#39;s file tree.
var Download = modfetch.Download

// DownloadDir returns the directory to which m should be downloaded.
// Note that the directory may not yet exist.
var DownloadDir = modfetch.DownloadDir

// DownloadZip downloads the specific module version to the
// local zip cache and returns the name of the zip file.
var DownloadZip = modfetch.DownloadZip

// GoMod is like Lookup(path).GoMod(rev) but avoids the
// repository path resolution in Lookup if the result is
// already cached on local disk.
var GoMod = modfetch.GoMod

// GoModFile is like GoMod but returns the name of the file containing
// the cached information.
var GoModFile = modfetch.GoModFile

// GoModSum returns the go.sum entry for the module version&#39;s go.mod file.
// (That is, it returns the entry listed in go.sum as &#34;path version/go.mod&#34;.)
var GoModSum = modfetch.GoModSum

// ImportRepoRev returns the module and version to use to access
// the given import path loaded from the source code repository that
// the original &#34;go get&#34; would have used, at the specific repository revision
// (typically a commit hash, but possibly also a source control tag).
var ImportRepoRev = modfetch.ImportRepoRev

// InfoFile is like Stat but returns the name of the file containing
// the cached information.
var InfoFile = modfetch.InfoFile

// IsPseudoVersion reports whether v is a pseudo-version.
var IsPseudoVersion = modfetch.IsPseudoVersion

// Lookup returns the module with the given module path.
// A successful return does not guarantee that the module
// has any defined versions.
var Lookup = modfetch.Lookup

// PseudoVersion returns a pseudo-version for the given major version (&#34;v1&#34;)
// preexisting older tagged version (&#34;&#34; or &#34;v1.2.3&#34; or &#34;v1.2.3-pre&#34;), revision time,
// and revision identifier (usually a 12-byte commit hash prefix).
var PseudoVersion = modfetch.PseudoVersion

// PseudoVersionRev returns the revision identifier of the pseudo-version v.
// It returns an error if v is not a pseudo-version.
var PseudoVersionRev = modfetch.PseudoVersionRev

// PseudoVersionTime returns the time stamp of the pseudo-version v.
// It returns an error if v is not a pseudo-version or if the time stamp
// embedded in the pseudo-version is not a valid time.
var PseudoVersionTime = modfetch.PseudoVersionTime

// RemoveAll removes a directory written by Download or Unzip, first applying
// any permission changes needed to do so.
var RemoveAll = modfetch.RemoveAll

// SideLock locks a file within the module cache that that guards edits to files
// outside the cache, such as go.sum and go.mod files in the user&#39;s working
// directory. It returns a function that must be called to unlock the file.
var SideLock = modfetch.SideLock

var SortVersions = modfetch.SortVersions

// Stat is like Lookup(path).Stat(rev) but avoids the
// repository path resolution in Lookup if the result is
// already cached on local disk.
var Stat = modfetch.Stat

// Sum returns the checksum for the downloaded copy of the given module,
// if present in the download cache.
var Sum = modfetch.Sum

// TrimGoSum trims go.sum to contain only the modules for which keep[m] is true.
var TrimGoSum = modfetch.TrimGoSum

var Unzip = modfetch.Unzip

// WriteGoSum writes the go.sum file if it needs to be updated.
var WriteGoSum = modfetch.WriteGoSum
