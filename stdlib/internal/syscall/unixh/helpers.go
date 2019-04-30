package unixh

import (
	"internal/syscall/unix"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	FstatatKey = "Fstatat"

	GetEntropyKey = "GetEntropy"

	GetRandomKey = "GetRandom"

	GetRandomKey = "GetRandom"

	IoctlKey = "Ioctl"

	IsNonblockKey = "IsNonblock"

	IsNonblockKey = "IsNonblock"

	IsNonblockKey = "IsNonblock"

	IsNonblockKey = "IsNonblock"

	OpenatKey = "Openat"

	OpenatKey = "Openat"

	OpenatKey = "Openat"

	OpenatKey = "Openat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"

	UnlinkatKey = "Unlinkat"
)

func New() hctx.Map {
	return hctx.Map{

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		FstatatKey: Fstatat,

		GetEntropyKey: GetEntropy,

		GetRandomKey: GetRandom,

		GetRandomKey: GetRandom,

		IoctlKey: Ioctl,

		IsNonblockKey: IsNonblock,

		IsNonblockKey: IsNonblock,

		IsNonblockKey: IsNonblock,

		IsNonblockKey: IsNonblock,

		OpenatKey: Openat,

		OpenatKey: Openat,

		OpenatKey: Openat,

		OpenatKey: Openat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,

		UnlinkatKey: Unlinkat,
	}
}

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

var Fstatat = unix.Fstatat

// GetEntropy calls the OpenBSD getentropy system call.
var GetEntropy = unix.GetEntropy

// GetRandom calls the FreeBSD getrandom system call.
var GetRandom = unix.GetRandom

// GetRandom calls the Linux getrandom system call.
// See https://git.kernel.org/cgit/linux/kernel/git/torvalds/linux.git/commit/?id=c6e9d6f38894798696f23c8084ca7edbf16ee895
var GetRandom = unix.GetRandom

var Ioctl = unix.Ioctl

var IsNonblock = unix.IsNonblock

var IsNonblock = unix.IsNonblock

var IsNonblock = unix.IsNonblock

var IsNonblock = unix.IsNonblock

var Openat = unix.Openat

var Openat = unix.Openat

var Openat = unix.Openat

var Openat = unix.Openat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat

var Unlinkat = unix.Unlinkat
