package geth

import (
	"cmd/go/internal/get"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	CheckImportPathKey = "CheckImportPath"

	RepoRootForImportPathKey = "RepoRootForImportPath"
)

func New() hctx.Map {
	return hctx.Map{

		CheckImportPathKey: CheckImportPath,

		RepoRootForImportPathKey: RepoRootForImportPath,
	}
}

// CheckImportPath checks that an import path is valid.
var CheckImportPath = get.CheckImportPath

// RepoRootForImportPath analyzes importPath to determine the
// version control system, and code repository to use.
var RepoRootForImportPath = get.RepoRootForImportPath
