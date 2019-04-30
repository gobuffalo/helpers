package loadh

import (
	"cmd/go/internal/load"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ClearCmdCacheKey = "ClearCmdCache"

	ClearPackageCacheKey = "ClearPackageCache"

	ClearPackageCachePartialKey = "ClearPackageCachePartial"

	FindVendorKey = "FindVendor"

	GoFilesPackageKey = "GoFilesPackage"

	ImportPathsKey = "ImportPaths"

	InstallTargetDirKey = "InstallTargetDir"

	LinkerDepsKey = "LinkerDeps"

	LoadImportKey = "LoadImport"

	LoadPackageKey = "LoadPackage"

	LoadPackageNoFlagsKey = "LoadPackageNoFlags"

	MatchPackageKey = "MatchPackage"

	ModuleImportPathKey = "ModuleImportPath"

	PackageListKey = "PackageList"

	PackagesKey = "Packages"

	PackagesAndErrorsKey = "PackagesAndErrors"

	PackagesForBuildKey = "PackagesForBuild"

	ReloadPackageNoFlagsKey = "ReloadPackageNoFlags"

	ResolveImportPathKey = "ResolveImportPath"

	SafeArgKey = "SafeArg"

	TestPackageListKey = "TestPackageList"

	TestPackagesForKey = "TestPackagesFor"

	VendoredImportPathKey = "VendoredImportPath"
)

func New() hctx.Map {
	return hctx.Map{

		ClearCmdCacheKey: ClearCmdCache,

		ClearPackageCacheKey: ClearPackageCache,

		ClearPackageCachePartialKey: ClearPackageCachePartial,

		FindVendorKey: FindVendor,

		GoFilesPackageKey: GoFilesPackage,

		ImportPathsKey: ImportPaths,

		InstallTargetDirKey: InstallTargetDir,

		LinkerDepsKey: LinkerDeps,

		LoadImportKey: LoadImport,

		LoadPackageKey: LoadPackage,

		LoadPackageNoFlagsKey: LoadPackageNoFlags,

		MatchPackageKey: MatchPackage,

		ModuleImportPathKey: ModuleImportPath,

		PackageListKey: PackageList,

		PackagesKey: Packages,

		PackagesAndErrorsKey: PackagesAndErrors,

		PackagesForBuildKey: PackagesForBuild,

		ReloadPackageNoFlagsKey: ReloadPackageNoFlags,

		ResolveImportPathKey: ResolveImportPath,

		SafeArgKey: SafeArg,

		TestPackageListKey: TestPackageList,

		TestPackagesForKey: TestPackagesFor,

		VendoredImportPathKey: VendoredImportPath,
	}
}

var ClearCmdCache = load.ClearCmdCache

var ClearPackageCache = load.ClearPackageCache

var ClearPackageCachePartial = load.ClearPackageCachePartial

// FindVendor looks for the last non-terminating &#34;vendor&#34; path element in the given import path.
// If there isn&#39;t one, FindVendor returns ok=false.
// Otherwise, FindVendor returns ok=true and the index of the &#34;vendor&#34;.
//
// Note that terminating &#34;vendor&#34; elements don&#39;t count: &#34;x/vendor&#34; is its own package,
// not the vendored copy of an import &#34;&#34; (the empty import path).
// This will allow people to have packages or commands named vendor.
// This may help reduce breakage, or it may just be confusing. We&#39;ll see.
var FindVendor = load.FindVendor

// GoFilesPackage creates a package for building a collection of Go files
// (typically named on the command line). The target is named p.a for
// package p or named after the first Go file for package main.
var GoFilesPackage = load.GoFilesPackage

var ImportPaths = load.ImportPaths

// InstallTargetDir reports the target directory for installing the command p.
var InstallTargetDir = load.InstallTargetDir

// LinkerDeps returns the list of linker-induced dependencies for main package p.
var LinkerDeps = load.LinkerDeps

// LoadImport scans the directory named by path, which must be an import path,
// but possibly a local import path (an absolute file system path or one beginning
// with ./ or ../). A local relative path is interpreted relative to srcDir.
// It returns a *Package describing the package found in that directory.
// LoadImport does not set tool flags and should only be used by
// this package, as part of a bigger load operation, and by GOPATH-based &#34;go get&#34;.
// TODO(rsc): When GOPATH-based &#34;go get&#34; is removed, unexport this function.
var LoadImport = load.LoadImport

// LoadPackage loads the package named by arg.
var LoadPackage = load.LoadPackage

// LoadPackageNoFlags is like LoadPackage
// but does not guarantee that the build tool flags are set in the result.
// It is only for use by GOPATH-based &#34;go get&#34;
// and is only appropriate for preliminary loading of packages.
// A real load using LoadPackage or (more likely)
// Packages, PackageAndErrors, or PackagesForBuild
// must be done before passing the package to any build
// steps, so that the tool flags can be set properly.
// TODO(rsc): When GOPATH-based &#34;go get&#34; is removed, delete this function.
var LoadPackageNoFlags = load.LoadPackageNoFlags

// MatchPackage(pattern, cwd)(p) reports whether package p matches pattern in the working directory cwd.
var MatchPackage = load.MatchPackage

// ModuleImportPath translates import paths found in go modules
// back down to paths that can be resolved in ordinary builds.
//
// Define “new” code as code with a go.mod file in the same directory
// or a parent directory. If an import in new code says x/y/v2/z but
// x/y/v2/z does not exist and x/y/go.mod says “module x/y/v2”,
// then go build will read the import as x/y/z instead.
// See golang.org/issue/25069.
var ModuleImportPath = load.ModuleImportPath

// PackageList returns the list of packages in the dag rooted at roots
// as visited in a depth-first post-order traversal.
var PackageList = load.PackageList

// Packages returns the packages named by the
// command line arguments &#39;args&#39;. If a named package
// cannot be loaded at all (for example, if the directory does not exist),
// then packages prints an error and does not include that
// package in the results. However, if errors occur trying
// to load dependencies of a named package, the named
// package is still returned, with p.Incomplete = true
// and details in p.DepsErrors.
var Packages = load.Packages

// PackagesAndErrors is like &#39;packages&#39; but returns a
// *Package for every argument, even the ones that
// cannot be loaded at all.
// The packages that fail to load will have p.Error != nil.
var PackagesAndErrors = load.PackagesAndErrors

// PackagesForBuild is like Packages but exits
// if any of the packages or their dependencies have errors
// (cannot be built).
var PackagesForBuild = load.PackagesForBuild

// ReloadPackageNoFlags is like LoadPackageNoFlags but makes sure
// not to use the package cache.
// It is only for use by GOPATH-based &#34;go get&#34;.
// TODO(rsc): When GOPATH-based &#34;go get&#34; is removed, delete this function.
var ReloadPackageNoFlags = load.ReloadPackageNoFlags

// ResolveImportPath returns the true meaning of path when it appears in parent.
// There are two different resolutions applied.
// First, there is Go 1.5 vendoring (golang.org/s/go15vendor).
// If vendor expansion doesn&#39;t trigger, then the path is also subject to
// Go 1.11 module legacy conversion (golang.org/issue/25069).
var ResolveImportPath = load.ResolveImportPath

// SafeArg reports whether arg is a &#34;safe&#34; command-line argument,
// meaning that when it appears in a command-line, it probably
// doesn&#39;t have some special meaning other than its own name.
// Obviously args beginning with - are not safe (they look like flags).
// Less obviously, args beginning with @ are not safe (they look like
// GNU binutils flagfile specifiers, sometimes called &#34;response files&#34;).
// To be conservative, we reject almost any arg beginning with non-alphanumeric ASCII.
// We accept leading . _ and / as likely in file system paths.
// There is a copy of this function in cmd/compile/internal/gc/noder.go.
var SafeArg = load.SafeArg

// TestPackageList returns the list of packages in the dag rooted at roots
// as visited in a depth-first post-order traversal, including the test
// imports of the roots. This ignores errors in test packages.
var TestPackageList = load.TestPackageList

// TestPackagesFor returns three packages:
// 	- ptest, the package p compiled with added &#34;package p&#34; test files.
// 	- pxtest, the result of compiling any &#34;package p_test&#34; (external) test files.
// 	- pmain, the package main corresponding to the test binary (running tests in ptest and pxtest).
//
// If the package has no &#34;package p_test&#34; test files, pxtest will be nil.
// If the non-test compilation of package p can be reused
// (for example, if there are no &#34;package p&#34; test files and
// package p need not be instrumented for coverage or any other reason),
// then the returned ptest == p.
//
// The caller is expected to have checked that len(p.TestGoFiles)+len(p.XTestGoFiles) &gt; 0,
// or else there&#39;s no point in any of this.
var TestPackagesFor = load.TestPackagesFor

// VendoredImportPath returns the vendor-expansion of path when it appears in parent.
// If parent is x/y/z, then path might expand to x/y/z/vendor/path, x/y/vendor/path,
// x/vendor/path, vendor/path, or else stay path if none of those exist.
// VendoredImportPath returns the expanded path or, if no expansion is found, the original.
var VendoredImportPath = load.VendoredImportPath
