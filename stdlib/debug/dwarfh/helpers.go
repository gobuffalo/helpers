package dwarfh

import (
	"debug/dwarf"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	NewKey = "New"
)

func New() hctx.Map {
	return hctx.Map{

		NewKey: New,
	}
}

// New returns a new Data object initialized from the given parameters.
// Rather than calling this function directly, clients should typically use
// the DWARF method of the File type of the appropriate package debug/elf,
// debug/macho, or debug/pe.
//
// The []byte arguments are the data from the corresponding debug section
// in the object file; for example, for an ELF object, abbrev is the contents of
// the &#34;.debug_abbrev&#34; section.
var New = dwarf.New
