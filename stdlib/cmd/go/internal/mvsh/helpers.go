package mvsh

import (
	"cmd/go/internal/mvs"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	BuildListKey = "BuildList"

	DowngradeKey = "Downgrade"

	ReqKey = "Req"

	UpgradeKey = "Upgrade"

	UpgradeAllKey = "UpgradeAll"
)

func New() hctx.Map {
	return hctx.Map{

		BuildListKey: BuildList,

		DowngradeKey: Downgrade,

		ReqKey: Req,

		UpgradeKey: Upgrade,

		UpgradeAllKey: UpgradeAll,
	}
}

// BuildList returns the build list for the target module.
// The first element is the target itself, with the remainder of the list sorted by path.
var BuildList = mvs.BuildList

// Downgrade returns a build list for the target module
// in which the given additional modules are downgraded.
//
// The versions to be downgraded may be unreachable from reqs.Latest and
// reqs.Previous, but the methods of reqs must otherwise handle such versions
// correctly.
var Downgrade = mvs.Downgrade

// Req returns the minimal requirement list for the target module
// that results in the given build list, with the constraint that all
// module paths listed in base must appear in the returned list.
var Req = mvs.Req

// Upgrade returns a build list for the target module
// in which the given additional modules are upgraded.
var Upgrade = mvs.Upgrade

// UpgradeAll returns a build list for the target module
// in which every module is upgraded to its latest version.
var UpgradeAll = mvs.UpgradeAll
