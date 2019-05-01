package userh

import (
	"os/user"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	CurrentKey = "Current"

	LookupKey = "Lookup"

	LookupGroupKey = "LookupGroup"

	LookupGroupIdKey = "LookupGroupId"

	LookupIdKey = "LookupId"
)

func New() hctx.Map {
	return hctx.Map{

		CurrentKey: Current,

		LookupKey: Lookup,

		LookupGroupKey: LookupGroup,

		LookupGroupIdKey: LookupGroupId,

		LookupIdKey: LookupId,
	}
}

// Current returns the current user.
var Current = user.Current

// Lookup looks up a user by username. If the user cannot be found, the
// returned error is of type UnknownUserError.
var Lookup = user.Lookup

// LookupGroup looks up a group by name. If the group cannot be found, the
// returned error is of type UnknownGroupError.
var LookupGroup = user.LookupGroup

// LookupGroupId looks up a group by groupid. If the group cannot be found, the
// returned error is of type UnknownGroupIdError.
var LookupGroupId = user.LookupGroupId

// LookupId looks up a user by userid. If the user cannot be found, the
// returned error is of type UnknownUserIdError.
var LookupId = user.LookupId
