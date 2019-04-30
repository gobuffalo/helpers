package registryh

import (
	"cmd/vendor/golang.org/x/sys/windows/registry"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	CreateKeyKey = "CreateKey"

	DeleteKeyKey = "DeleteKey"

	ExpandStringKey = "ExpandString"

	LoadRegLoadMUIStringKey = "LoadRegLoadMUIString"

	OpenKeyKey = "OpenKey"

	OpenRemoteKeyKey = "OpenRemoteKey"
)

func New() hctx.Map {
	return hctx.Map{

		CreateKeyKey: CreateKey,

		DeleteKeyKey: DeleteKey,

		ExpandStringKey: ExpandString,

		LoadRegLoadMUIStringKey: LoadRegLoadMUIString,

		OpenKeyKey: OpenKey,

		OpenRemoteKeyKey: OpenRemoteKey,
	}
}

// CreateKey creates a key named path under open key k.
// CreateKey returns the new key and a boolean flag that reports
// whether the key already existed.
// The access parameter specifies the access rights for the key
// to be created.
var CreateKey = registry.CreateKey

// DeleteKey deletes the subkey path of key k and its values.
var DeleteKey = registry.DeleteKey

// ExpandString expands environment-variable strings and replaces
// them with the values defined for the current user.
// Use ExpandString to expand EXPAND_SZ strings.
var ExpandString = registry.ExpandString

var LoadRegLoadMUIString = registry.LoadRegLoadMUIString

// OpenKey opens a new key with path name relative to key k.
// It accepts any open key, including CURRENT_USER and others,
// and returns the new key and an error.
// The access parameter specifies desired access rights to the
// key to be opened.
var OpenKey = registry.OpenKey

// OpenRemoteKey opens a predefined registry key on another
// computer pcname. The key to be opened is specified by k, but
// can only be one of LOCAL_MACHINE, PERFORMANCE_DATA or USERS.
// If pcname is &#34;&#34;, OpenRemoteKey returns local computer key.
var OpenRemoteKey = registry.OpenRemoteKey
