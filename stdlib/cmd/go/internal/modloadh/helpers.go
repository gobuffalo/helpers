package modloadh

import (
	"cmd/go/internal/modload"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AllowWriteGoModKey = "AllowWriteGoMod"

	AllowedKey = "Allowed"

	BinDirKey = "BinDir"

	BuildListKey = "BuildList"

	DirImportPathKey = "DirImportPath"

	DisallowWriteGoModKey = "DisallowWriteGoMod"

	EnabledKey = "Enabled"

	FindModulePathKey = "FindModulePath"

	FindModuleRootKey = "FindModuleRoot"

	HasModRootKey = "HasModRoot"

	ImportKey = "Import"

	ImportFromFilesKey = "ImportFromFiles"

	ImportMapKey = "ImportMap"

	ImportPathsKey = "ImportPaths"

	InitKey = "Init"

	InitGoStmtKey = "InitGoStmt"

	InitModKey = "InitMod"

	ListModulesKey = "ListModules"

	LoadALLKey = "LoadALL"

	LoadBuildListKey = "LoadBuildList"

	LoadVendorKey = "LoadVendor"

	LookupKey = "Lookup"

	MinReqsKey = "MinReqs"

	ModFileKey = "ModFile"

	ModInfoProgKey = "ModInfoProg"

	ModRootKey = "ModRoot"

	ModuleInfoKey = "ModuleInfo"

	ModuleUsedDirectlyKey = "ModuleUsedDirectly"

	PackageBuildInfoKey = "PackageBuildInfo"

	PackageDirKey = "PackageDir"

	PackageModuleKey = "PackageModule"

	PackageModuleInfoKey = "PackageModuleInfo"

	QueryKey = "Query"

	QueryPackageKey = "QueryPackage"

	ReloadBuildListKey = "ReloadBuildList"

	ReplacementKey = "Replacement"

	ReqsKey = "Reqs"

	SetBuildListKey = "SetBuildList"

	TargetPackagesKey = "TargetPackages"

	WhyKey = "Why"

	WhyDepthKey = "WhyDepth"

	WriteGoModKey = "WriteGoMod"
)

func New() hctx.Map {
	return hctx.Map{

		AllowWriteGoModKey: AllowWriteGoMod,

		AllowedKey: Allowed,

		BinDirKey: BinDir,

		BuildListKey: BuildList,

		DirImportPathKey: DirImportPath,

		DisallowWriteGoModKey: DisallowWriteGoMod,

		EnabledKey: Enabled,

		FindModulePathKey: FindModulePath,

		FindModuleRootKey: FindModuleRoot,

		HasModRootKey: HasModRoot,

		ImportKey: Import,

		ImportFromFilesKey: ImportFromFiles,

		ImportMapKey: ImportMap,

		ImportPathsKey: ImportPaths,

		InitKey: Init,

		InitGoStmtKey: InitGoStmt,

		InitModKey: InitMod,

		ListModulesKey: ListModules,

		LoadALLKey: LoadALL,

		LoadBuildListKey: LoadBuildList,

		LoadVendorKey: LoadVendor,

		LookupKey: Lookup,

		MinReqsKey: MinReqs,

		ModFileKey: ModFile,

		ModInfoProgKey: ModInfoProg,

		ModRootKey: ModRoot,

		ModuleInfoKey: ModuleInfo,

		ModuleUsedDirectlyKey: ModuleUsedDirectly,

		PackageBuildInfoKey: PackageBuildInfo,

		PackageDirKey: PackageDir,

		PackageModuleKey: PackageModule,

		PackageModuleInfoKey: PackageModuleInfo,

		QueryKey: Query,

		QueryPackageKey: QueryPackage,

		ReloadBuildListKey: ReloadBuildList,

		ReplacementKey: Replacement,

		ReqsKey: Reqs,

		SetBuildListKey: SetBuildList,

		TargetPackagesKey: TargetPackages,

		WhyKey: Why,

		WhyDepthKey: WhyDepth,

		WriteGoModKey: WriteGoMod,
	}
}

// AllowWriteGoMod undoes the effect of DisallowWriteGoMod:
// future calls to WriteGoMod will update go.mod if needed.
// Note that any past calls have been discarded, so typically
// a call to AlowWriteGoMod should be followed by a call to WriteGoMod.
var AllowWriteGoMod = modload.AllowWriteGoMod

// Allowed reports whether module m is allowed (not excluded) by the main module&#39;s go.mod.
var Allowed = modload.Allowed

var BinDir = modload.BinDir

// BuildList returns the module build list,
// typically constructed by a previous call to
// LoadBuildList or ImportPaths.
// The caller must not modify the returned list.
var BuildList = modload.BuildList

// DirImportPath returns the effective import path for dir,
// provided it is within the main module, or else returns &#34;.&#34;.
var DirImportPath = modload.DirImportPath

// DisallowWriteGoMod causes future calls to WriteGoMod to do nothing at all.
var DisallowWriteGoMod = modload.DisallowWriteGoMod

// Enabled reports whether modules are (or must be) enabled.
// If modules are enabled but there is no main module, Enabled returns true
// and then the first use of module information will call die
// (usually through MustModRoot).
var Enabled = modload.Enabled

// Exported only for testing.
var FindModulePath = modload.FindModulePath

// Exported only for testing.
var FindModuleRoot = modload.FindModuleRoot

// HasModRoot reports whether a main module is present.
// HasModRoot may return false even if Enabled returns true: for example, &#39;get&#39;
// does not require a main module.
var HasModRoot = modload.HasModRoot

// Import finds the module and directory in the build list
// containing the package with the given import path.
// The answer must be unique: Import returns an error
// if multiple modules attempt to provide the same package.
// Import can return a module with an empty m.Path, for packages in the standard library.
// Import can return an empty directory string, for fake packages like &#34;C&#34; and &#34;unsafe&#34;.
//
// If the package cannot be found in the current build list,
// Import returns an ImportMissingError as the error.
// If Import can identify a module that could be added to supply the package,
// the ImportMissingError records that module.
var Import = modload.Import

// ImportFromFiles adds modules to the build list as needed
// to satisfy the imports in the named Go source files.
var ImportFromFiles = modload.ImportFromFiles

// ImportMap returns the actual package import path
// for an import path found in source code.
// If the given import path does not appear in the source code
// for the packages that have been loaded, ImportMap returns the empty string.
var ImportMap = modload.ImportMap

// ImportPaths returns the set of packages matching the args (patterns),
// adding modules to the build list as needed to satisfy new imports.
var ImportPaths = modload.ImportPaths

// Init determines whether module mode is enabled, locates the root of the
// current module (if any), sets environment variables for Git subprocesses, and
// configures the cfg, codehost, load, modfetch, and search packages for use
// with modules.
var Init = modload.Init

// InitGoStmt adds a go statement, unless there already is one.
var InitGoStmt = modload.InitGoStmt

// InitMod sets Target and, if there is a main module, parses the initial build
// list from its go.mod file, creating and populating that file if needed.
var InitMod = modload.InitMod

var ListModules = modload.ListModules

// LoadALL returns the set of all packages in the current module
// and their dependencies in any other modules, without filtering
// due to build tags, except &#34;+build ignore&#34;.
// It adds modules to the build list as needed to satisfy new imports.
// This set is useful for deciding whether a particular import is needed
// anywhere in a module.
var LoadALL = modload.LoadALL

// LoadBuildList loads and returns the build list from go.mod.
// The loading of the build list happens automatically in ImportPaths:
// LoadBuildList need only be called if ImportPaths is not
// (typically in commands that care about the module but
// no particular package).
var LoadBuildList = modload.LoadBuildList

// LoadVendor is like LoadALL but only follows test dependencies
// for tests in the main module. Tests in dependency modules are
// ignored completely.
// This set is useful for identifying the which packages to include in a vendor directory.
var LoadVendor = modload.LoadVendor

// Lookup returns the source directory, import path, and any loading error for
// the package at path.
// Lookup requires that one of the Load functions in this package has already
// been called.
var Lookup = modload.Lookup

// MinReqs returns a Reqs with minimal dependencies of Target,
// as will be written to go.mod.
var MinReqs = modload.MinReqs

// ModFile returns the parsed go.mod file.
//
// Note that after calling ImportPaths or LoadBuildList,
// the require statements in the modfile.File are no longer
// the source of truth and will be ignored: edits made directly
// will be lost at the next call to WriteGoMod.
// To make permanent changes to the require statements
// in go.mod, edit it before calling ImportPaths or LoadBuildList.
var ModFile = modload.ModFile

var ModInfoProg = modload.ModInfoProg

// ModRoot returns the root of the main module.
// It calls base.Fatalf if there is no main module.
var ModRoot = modload.ModRoot

var ModuleInfo = modload.ModuleInfo

// ModuleUsedDirectly reports whether the main module directly imports
// some package in the module with the given path.
var ModuleUsedDirectly = modload.ModuleUsedDirectly

var PackageBuildInfo = modload.PackageBuildInfo

// PackageDir returns the directory containing the source code
// for the package named by the import path.
var PackageDir = modload.PackageDir

// PackageModule returns the module providing the package named by the import path.
var PackageModule = modload.PackageModule

var PackageModuleInfo = modload.PackageModuleInfo

// Query looks up a revision of a given module given a version query string.
// The module must be a complete module path.
// The version must take one of the following forms:
//
// 	- the literal string &#34;latest&#34;, denoting the latest available, allowed tagged version,
// 	  with non-prereleases preferred over prereleases.
// 	  If there are no tagged versions in the repo, latest returns the most recent commit.
// 	- v1, denoting the latest available tagged version v1.x.x.
// 	- v1.2, denoting the latest available tagged version v1.2.x.
// 	- v1.2.3, a semantic version string denoting that tagged version.
// 	- &lt;v1.2.3, &lt;=v1.2.3, &gt;v1.2.3, &gt;=v1.2.3,
// 	   denoting the version closest to the target and satisfying the given operator,
// 	   with non-prereleases preferred over prereleases.
// 	- a repository commit identifier, denoting that commit.
//
// If the allowed function is non-nil, Query excludes any versions for which allowed returns false.
//
// If path is the path of the main module and the query is &#34;latest&#34;,
// Query returns Target.Version as the version.
var Query = modload.Query

// QueryPackage looks up a revision of a module containing path.
//
// If multiple modules with revisions matching the query provide the requested
// package, QueryPackage picks the one with the longest module path.
//
// If the path is in the main module and the query is &#34;latest&#34;,
// QueryPackage returns Target as the version.
var QueryPackage = modload.QueryPackage

var ReloadBuildList = modload.ReloadBuildList

// Replacement returns the replacement for mod, if any, from go.mod.
// If there is no replacement for mod, Replacement returns
// a module.Version with Path == &#34;&#34;.
var Replacement = modload.Replacement

// Reqs returns the current module requirement graph.
// Future calls to SetBuildList do not affect the operation
// of the returned Reqs.
var Reqs = modload.Reqs

// SetBuildList sets the module build list.
// The caller is responsible for ensuring that the list is valid.
// SetBuildList does not retain a reference to the original list.
var SetBuildList = modload.SetBuildList

// TargetPackages returns the list of packages in the target (top-level) module,
// under all build tag settings.
var TargetPackages = modload.TargetPackages

// Why returns the &#34;go mod why&#34; output stanza for the given package,
// without the leading # comment.
// The package graph must have been loaded already, usually by LoadALL.
// If there is no reason for the package to be in the current build,
// Why returns an empty string.
var Why = modload.Why

// WhyDepth returns the number of steps in the Why listing.
// If there is no reason for the package to be in the current build,
// WhyDepth returns 0.
var WhyDepth = modload.WhyDepth

// WriteGoMod writes the current build list back to go.mod.
var WriteGoMod = modload.WriteGoMod
