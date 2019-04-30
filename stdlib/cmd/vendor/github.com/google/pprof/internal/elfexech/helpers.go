package elfexech

import (
	"cmd/vendor/github.com/google/pprof/internal/elfexec"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	FindTextProgHeaderKey = "FindTextProgHeader"

	GetBaseKey = "GetBase"

	GetBuildIDKey = "GetBuildID"
)

func New() hctx.Map {
	return hctx.Map{

		FindTextProgHeaderKey: FindTextProgHeader,

		GetBaseKey: GetBase,

		GetBuildIDKey: GetBuildID,
	}
}

// FindTextProgHeader finds the program segment header containing the .text
// section or nil if the segment cannot be found.
var FindTextProgHeader = elfexec.FindTextProgHeader

// GetBase determines the base address to subtract from virtual
// address to get symbol table address. For an executable, the base
// is 0. Otherwise, it&#39;s a shared library, and the base is the
// address where the mapping starts. The kernel is special, and may
// use the address of the _stext symbol as the mmap start. _stext
// offset can be obtained with `nm vmlinux | grep _stext`
var GetBase = elfexec.GetBase

// GetBuildID returns the GNU build-ID for an ELF binary.
//
// If no build-ID was found but the binary was read without error, it returns
// (nil, nil).
var GetBuildID = elfexec.GetBuildID
