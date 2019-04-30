package objabih

import (
	"cmd/internal/objabi"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AbsFileKey = "AbsFile"

	AddVersionFlagKey = "AddVersionFlag"

	DefaultExpstringKey = "DefaultExpstring"

	ExpstringKey = "Expstring"

	FlagcountKey = "Flagcount"

	Flagfn1Key = "Flagfn1"

	FlagparseKey = "Flagparse"

	FlagprintKey = "Flagprint"

	Framepointer_enabledKey = "Framepointer_enabled"

	GetFuncIDKey = "GetFuncID"

	GetgoextlinkenabledKey = "Getgoextlinkenabled"

	PathToPrefixKey = "PathToPrefix"

	WorkingDirKey = "WorkingDir"
)

func New() hctx.Map {
	return hctx.Map{

		AbsFileKey: AbsFile,

		AddVersionFlagKey: AddVersionFlag,

		DefaultExpstringKey: DefaultExpstring,

		ExpstringKey: Expstring,

		FlagcountKey: Flagcount,

		Flagfn1Key: Flagfn1,

		FlagparseKey: Flagparse,

		FlagprintKey: Flagprint,

		Framepointer_enabledKey: Framepointer_enabled,

		GetFuncIDKey: GetFuncID,

		GetgoextlinkenabledKey: Getgoextlinkenabled,

		PathToPrefixKey: PathToPrefix,

		WorkingDirKey: WorkingDir,
	}
}

// AbsFile returns the absolute filename for file in the given directory.
// It also removes a leading pathPrefix, or else rewrites a leading $GOROOT
// prefix to the literal &#34;$GOROOT&#34;.
// If the resulting path is the empty string, the result is &#34;??&#34;.
var AbsFile = objabi.AbsFile

var AddVersionFlag = objabi.AddVersionFlag

var DefaultExpstring = objabi.DefaultExpstring

var Expstring = objabi.Expstring

var Flagcount = objabi.Flagcount

var Flagfn1 = objabi.Flagfn1

var Flagparse = objabi.Flagparse

var Flagprint = objabi.Flagprint

var Framepointer_enabled = objabi.Framepointer_enabled

// Get the function ID for the named function in the named file.
// The function should be package-qualified.
var GetFuncID = objabi.GetFuncID

var Getgoextlinkenabled = objabi.Getgoextlinkenabled

// PathToPrefix converts raw string to the prefix that will be used in the
// symbol table. All control characters, space, &#39;%&#39; and &#39;&#34;&#39;, as well as
// non-7-bit clean bytes turn into %xx. The period needs escaping only in the
// last segment of the path, and it makes for happier users if we escape that as
// little as possible.
var PathToPrefix = objabi.PathToPrefix

// WorkingDir returns the current working directory
// (or &#34;/???&#34; if the directory cannot be identified),
// with &#34;/&#34; as separator.
var WorkingDir = objabi.WorkingDir
