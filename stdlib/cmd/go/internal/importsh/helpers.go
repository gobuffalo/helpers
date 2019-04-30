package importsh

import (
	"cmd/go/internal/imports"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	MatchFileKey = "MatchFile"

	ReadCommentsKey = "ReadComments"

	ReadImportsKey = "ReadImports"

	ScanDirKey = "ScanDir"

	ScanFilesKey = "ScanFiles"

	ShouldBuildKey = "ShouldBuild"

	TagsKey = "Tags"
)

func New() hctx.Map {
	return hctx.Map{

		MatchFileKey: MatchFile,

		ReadCommentsKey: ReadComments,

		ReadImportsKey: ReadImports,

		ScanDirKey: ScanDir,

		ScanFilesKey: ScanFiles,

		ShouldBuildKey: ShouldBuild,

		TagsKey: Tags,
	}
}

// MatchFile returns false if the name contains a $GOOS or $GOARCH
// suffix which does not match the current system.
// The recognized name formats are:
//
//     name_$(GOOS).*
//     name_$(GOARCH).*
//     name_$(GOOS)_$(GOARCH).*
//     name_$(GOOS)_test.*
//     name_$(GOARCH)_test.*
//     name_$(GOOS)_$(GOARCH)_test.*
//
// An exception: if GOOS=android, then files with GOOS=linux are also matched.
//
// If tags[&#34;*&#34;] is true, then MatchFile will consider all possible
// GOOS and GOARCH to be available and will consequently
// always return true.
var MatchFile = imports.MatchFile

// ReadComments is like ioutil.ReadAll, except that it only reads the leading
// block of comments in the file.
var ReadComments = imports.ReadComments

// ReadImports is like ioutil.ReadAll, except that it expects a Go file as input
// and stops reading the input once the imports have completed.
var ReadImports = imports.ReadImports

var ScanDir = imports.ScanDir

var ScanFiles = imports.ScanFiles

// ShouldBuild reports whether it is okay to use this file,
// The rule is that in the file&#39;s leading run of // comments
// and blank lines, which must be followed by a blank line
// (to avoid including a Go package clause doc comment),
// lines beginning with &#39;// +build&#39; are taken as build directives.
//
// The file is accepted only if each such line lists something
// matching the file. For example:
//
// 	// +build windows linux
//
// marks the file as applicable only on Windows and Linux.
//
// If tags[&#34;*&#34;] is true, then ShouldBuild will consider every
// build tag except &#34;ignore&#34; to be both true and false for
// the purpose of satisfying build tags, in order to estimate
// (conservatively) whether a file could ever possibly be used
// in any build.
var ShouldBuild = imports.ShouldBuild

var Tags = imports.Tags
