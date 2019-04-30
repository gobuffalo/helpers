package armasmh

import (
	"cmd/vendor/golang.org/x/arch/arm/armasm"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	DecodeKey = "Decode"

	GNUSyntaxKey = "GNUSyntax"

	GoSyntaxKey = "GoSyntax"
)

func New() hctx.Map {
	return hctx.Map{

		DecodeKey: Decode,

		GNUSyntaxKey: GNUSyntax,

		GoSyntaxKey: GoSyntax,
	}
}

// Decode decodes the leading bytes in src as a single instruction.
var Decode = armasm.Decode

// GNUSyntax returns the GNU assembler syntax for the instruction, as defined by GNU binutils.
// This form typically matches the syntax defined in the ARM Reference Manual.
var GNUSyntax = armasm.GNUSyntax

// GoSyntax returns the Go assembler syntax for the instruction.
// The syntax was originally defined by Plan 9.
// The pc is the program counter of the instruction, used for expanding
// PC-relative addresses into absolute ones.
// The symname function queries the symbol table for the program
// being disassembled. Given a target address it returns the name and base
// address of the symbol containing the target, if any; otherwise it returns &#34;&#34;, 0.
// The reader r should read from the text segment using text addresses
// as offsets; it is used to display pc-relative loads as constant loads.
var GoSyntax = armasm.GoSyntax
