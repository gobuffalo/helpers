package x86asmh

import (
	"cmd/vendor/golang.org/x/arch/x86/x86asm"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	DecodeKey = "Decode"

	GNUSyntaxKey = "GNUSyntax"

	GoSyntaxKey = "GoSyntax"

	IntelSyntaxKey = "IntelSyntax"
)

func New() hctx.Map {
	return hctx.Map{

		DecodeKey: Decode,

		GNUSyntaxKey: GNUSyntax,

		GoSyntaxKey: GoSyntax,

		IntelSyntaxKey: IntelSyntax,
	}
}

// Decode decodes the leading bytes in src as a single instruction.
// The mode arguments specifies the assumed processor mode:
// 16, 32, or 64 for 16-, 32-, and 64-bit execution modes.
var Decode = x86asm.Decode

// GNUSyntax returns the GNU assembler syntax for the instruction, as defined by GNU binutils.
// This general form is often called ``AT&amp;T syntax&#39;&#39; as a reference to AT&amp;T System V Unix.
var GNUSyntax = x86asm.GNUSyntax

// GoSyntax returns the Go assembler syntax for the instruction.
// The syntax was originally defined by Plan 9.
// The pc is the program counter of the instruction, used for expanding
// PC-relative addresses into absolute ones.
// The symname function queries the symbol table for the program
// being disassembled. Given a target address it returns the name and base
// address of the symbol containing the target, if any; otherwise it returns &#34;&#34;, 0.
var GoSyntax = x86asm.GoSyntax

// IntelSyntax returns the Intel assembler syntax for the instruction, as defined by Intel&#39;s XED tool.
var IntelSyntax = x86asm.IntelSyntax
