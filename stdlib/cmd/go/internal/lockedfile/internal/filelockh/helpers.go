package filelockh

import (
	"cmd/go/internal/lockedfile/internal/filelock"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	IsNotSupportedKey = "IsNotSupported"

	LockKey = "Lock"

	RLockKey = "RLock"

	UnlockKey = "Unlock"
)

func New() hctx.Map {
	return hctx.Map{

		IsNotSupportedKey: IsNotSupported,

		LockKey: Lock,

		RLockKey: RLock,

		UnlockKey: Unlock,
	}
}

// IsNotSupported returns a boolean indicating whether the error is known to
// report that a function is not supported (possibly for a specific input).
// It is satisfied by ErrNotSupported as well as some syscall errors.
var IsNotSupported = filelock.IsNotSupported

// Lock places an advisory write lock on the file, blocking until it can be
// locked.
//
// If Lock returns nil, no other process will be able to place a read or write
// lock on the file until this process exits, closes f, or calls Unlock on it.
//
// If f&#39;s descriptor is already read- or write-locked, the behavior of Lock is
// unspecified.
//
// Closing the file may or may not release the lock promptly. Callers should
// ensure that Unlock is always called when Lock succeeds.
var Lock = filelock.Lock

// RLock places an advisory read lock on the file, blocking until it can be locked.
//
// If RLock returns nil, no other process will be able to place a write lock on
// the file until this process exits, closes f, or calls Unlock on it.
//
// If f is already read- or write-locked, the behavior of RLock is unspecified.
//
// Closing the file may or may not release the lock promptly. Callers should
// ensure that Unlock is always called if RLock succeeds.
var RLock = filelock.RLock

// Unlock removes an advisory lock placed on f by this process.
//
// The caller must not attempt to unlock a file that is not locked.
var Unlock = filelock.Unlock
