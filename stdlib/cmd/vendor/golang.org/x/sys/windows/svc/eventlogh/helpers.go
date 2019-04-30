package eventlogh

import (
	"cmd/vendor/golang.org/x/sys/windows/svc/eventlog"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	InstallKey = "Install"

	InstallAsEventCreateKey = "InstallAsEventCreate"

	OpenKey = "Open"

	OpenRemoteKey = "OpenRemote"

	RemoveKey = "Remove"
)

func New() hctx.Map {
	return hctx.Map{

		InstallKey: Install,

		InstallAsEventCreateKey: InstallAsEventCreate,

		OpenKey: Open,

		OpenRemoteKey: OpenRemote,

		RemoveKey: Remove,
	}
}

// Install modifies PC registry to allow logging with an event source src.
// It adds all required keys and values to the event log registry key.
// Install uses msgFile as the event message file. If useExpandKey is true,
// the event message file is installed as REG_EXPAND_SZ value,
// otherwise as REG_SZ. Use bitwise of log.Error, log.Warning and
// log.Info to specify events supported by the new event source.
var Install = eventlog.Install

// InstallAsEventCreate is the same as Install, but uses
// %SystemRoot%\System32\EventCreate.exe as the event message file.
var InstallAsEventCreate = eventlog.InstallAsEventCreate

// Open retrieves a handle to the specified event log.
var Open = eventlog.Open

// OpenRemote does the same as Open, but on different computer host.
var OpenRemote = eventlog.OpenRemote

// Remove deletes all registry elements installed by the correspondent Install.
var Remove = eventlog.Remove
