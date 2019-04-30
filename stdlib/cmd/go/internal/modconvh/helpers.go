package modconvh

import (
	"cmd/go/internal/modconv"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ConvertLegacyConfigKey = "ConvertLegacyConfig"

	ParseDependenciesTSVKey = "ParseDependenciesTSV"

	ParseGLOCKFILEKey = "ParseGLOCKFILE"

	ParseGlideLockKey = "ParseGlideLock"

	ParseGodepsJSONKey = "ParseGodepsJSON"

	ParseGopkgLockKey = "ParseGopkgLock"

	ParseVendorConfKey = "ParseVendorConf"

	ParseVendorJSONKey = "ParseVendorJSON"

	ParseVendorManifestKey = "ParseVendorManifest"

	ParseVendorYMLKey = "ParseVendorYML"
)

func New() hctx.Map {
	return hctx.Map{

		ConvertLegacyConfigKey: ConvertLegacyConfig,

		ParseDependenciesTSVKey: ParseDependenciesTSV,

		ParseGLOCKFILEKey: ParseGLOCKFILE,

		ParseGlideLockKey: ParseGlideLock,

		ParseGodepsJSONKey: ParseGodepsJSON,

		ParseGopkgLockKey: ParseGopkgLock,

		ParseVendorConfKey: ParseVendorConf,

		ParseVendorJSONKey: ParseVendorJSON,

		ParseVendorManifestKey: ParseVendorManifest,

		ParseVendorYMLKey: ParseVendorYML,
	}
}

// ConvertLegacyConfig converts legacy config to modfile.
// The file argument is slash-delimited.
var ConvertLegacyConfig = modconv.ConvertLegacyConfig

var ParseDependenciesTSV = modconv.ParseDependenciesTSV

var ParseGLOCKFILE = modconv.ParseGLOCKFILE

var ParseGlideLock = modconv.ParseGlideLock

var ParseGodepsJSON = modconv.ParseGodepsJSON

var ParseGopkgLock = modconv.ParseGopkgLock

var ParseVendorConf = modconv.ParseVendorConf

var ParseVendorJSON = modconv.ParseVendorJSON

var ParseVendorManifest = modconv.ParseVendorManifest

var ParseVendorYML = modconv.ParseVendorYML
