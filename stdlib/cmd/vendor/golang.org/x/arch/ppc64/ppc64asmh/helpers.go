package ppc64asmh

import (
	"cmd/vendor/golang.org/x/arch/ppc64/ppc64asm"
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

// Decode decodes the leading bytes in src as a single instruction using
// byte order ord.
var Decode = ppc64asm.Decode

// GNUSyntax returns the GNU assembler syntax for the instruction, as defined by GNU binutils.
// This form typically matches the syntax defined in the Power ISA Reference Manual.
var GNUSyntax = ppc64asm.GNUSyntax

// GoSyntax returns the Go assembler syntax for the instruction.
// The pc is the program counter of the first instruction, used for expanding
// PC-relative addresses into absolute ones.
// The symname function queries the symbol table for the program
// being disassembled. It returns the name and base address of the symbol
// containing the target, if any; otherwise it returns &#34;&#34;, 0.
var GoSyntax = ppc64asm.GoSyntax
