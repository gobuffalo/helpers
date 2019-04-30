package elfh

import (
	"debug/elf"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewFileKey = "NewFile"

	OpenKey = "Open"

	R_INFOKey = "R_INFO"

	R_INFO32Key = "R_INFO32"

	R_SYM32Key = "R_SYM32"

	R_SYM64Key = "R_SYM64"

	R_TYPE32Key = "R_TYPE32"

	R_TYPE64Key = "R_TYPE64"

	ST_BINDKey = "ST_BIND"

	ST_INFOKey = "ST_INFO"

	ST_TYPEKey = "ST_TYPE"

	ST_VISIBILITYKey = "ST_VISIBILITY"
)

func New() hctx.Map {
	return hctx.Map{

		NewFileKey: NewFile,

		OpenKey: Open,

		R_INFOKey: R_INFO,

		R_INFO32Key: R_INFO32,

		R_SYM32Key: R_SYM32,

		R_SYM64Key: R_SYM64,

		R_TYPE32Key: R_TYPE32,

		R_TYPE64Key: R_TYPE64,

		ST_BINDKey: ST_BIND,

		ST_INFOKey: ST_INFO,

		ST_TYPEKey: ST_TYPE,

		ST_VISIBILITYKey: ST_VISIBILITY,
	}
}

// NewFile creates a new File for accessing an ELF binary in an underlying reader.
// The ELF binary is expected to start at position 0 in the ReaderAt.
var NewFile = elf.NewFile

// Open opens the named file using os.Open and prepares it for use as an ELF binary.
var Open = elf.Open

var R_INFO = elf.R_INFO

var R_INFO32 = elf.R_INFO32

var R_SYM32 = elf.R_SYM32

var R_SYM64 = elf.R_SYM64

var R_TYPE32 = elf.R_TYPE32

var R_TYPE64 = elf.R_TYPE64

var ST_BIND = elf.ST_BIND

var ST_INFO = elf.ST_INFO

var ST_TYPE = elf.ST_TYPE

var ST_VISIBILITY = elf.ST_VISIBILITY
