package ldh

import (
	"cmd/link/internal/ld"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AdddynsymKey = "Adddynsym"

	AddstringKey = "Addstring"

	AsmbelfKey = "Asmbelf"

	AsmbelfsetupKey = "Asmbelfsetup"

	AsmbmachoKey = "Asmbmacho"

	AsmbpeKey = "Asmbpe"

	AsmbxcoffKey = "Asmbxcoff"

	AsmelfsymKey = "Asmelfsym"

	Asmplan9symKey = "Asmplan9sym"

	AtExitKey = "AtExit"

	CodeblkKey = "Codeblk"

	CodeblkPadKey = "CodeblkPad"

	CputimeKey = "Cputime"

	DatblkKey = "Datblk"

	DomacholinkKey = "Domacholink"

	DwarfblkKey = "Dwarfblk"

	ELF32_R_INFOKey = "ELF32_R_INFO"

	ELF32_R_SYMKey = "ELF32_R_SYM"

	ELF32_R_TYPEKey = "ELF32_R_TYPE"

	ELF32_ST_BINDKey = "ELF32_ST_BIND"

	ELF32_ST_INFOKey = "ELF32_ST_INFO"

	ELF32_ST_TYPEKey = "ELF32_ST_TYPE"

	ELF32_ST_VISIBILITYKey = "ELF32_ST_VISIBILITY"

	ELF64_R_INFOKey = "ELF64_R_INFO"

	ELF64_R_SYMKey = "ELF64_R_SYM"

	ELF64_R_TYPEKey = "ELF64_R_TYPE"

	ELF64_ST_BINDKey = "ELF64_ST_BIND"

	ELF64_ST_INFOKey = "ELF64_ST_INFO"

	ELF64_ST_TYPEKey = "ELF64_ST_TYPE"

	ELF64_ST_VISIBILITYKey = "ELF64_ST_VISIBILITY"

	ElfemitrelocKey = "Elfemitreloc"

	ElfinitKey = "Elfinit"

	ElfwritedynentKey = "Elfwritedynent"

	ElfwritedynentsymplusKey = "Elfwritedynentsymplus"

	EntryvalueKey = "Entryvalue"

	ErrorfKey = "Errorf"

	ExitKey = "Exit"

	ExitfKey = "Exitf"

	LflagKey = "Lflag"

	LoaderblkKey = "Loaderblk"

	MachoemitrelocKey = "Machoemitreloc"

	MainKey = "Main"

	PeinitKey = "Peinit"

	RndKey = "Rnd"

	SymaddrKey = "Symaddr"

	XcoffadddynrelKey = "Xcoffadddynrel"

	XcoffinitKey = "Xcoffinit"
)

func New() hctx.Map {
	return hctx.Map{

		AdddynsymKey: Adddynsym,

		AddstringKey: Addstring,

		AsmbelfKey: Asmbelf,

		AsmbelfsetupKey: Asmbelfsetup,

		AsmbmachoKey: Asmbmacho,

		AsmbpeKey: Asmbpe,

		AsmbxcoffKey: Asmbxcoff,

		AsmelfsymKey: Asmelfsym,

		Asmplan9symKey: Asmplan9sym,

		AtExitKey: AtExit,

		CodeblkKey: Codeblk,

		CodeblkPadKey: CodeblkPad,

		CputimeKey: Cputime,

		DatblkKey: Datblk,

		DomacholinkKey: Domacholink,

		DwarfblkKey: Dwarfblk,

		ELF32_R_INFOKey: ELF32_R_INFO,

		ELF32_R_SYMKey: ELF32_R_SYM,

		ELF32_R_TYPEKey: ELF32_R_TYPE,

		ELF32_ST_BINDKey: ELF32_ST_BIND,

		ELF32_ST_INFOKey: ELF32_ST_INFO,

		ELF32_ST_TYPEKey: ELF32_ST_TYPE,

		ELF32_ST_VISIBILITYKey: ELF32_ST_VISIBILITY,

		ELF64_R_INFOKey: ELF64_R_INFO,

		ELF64_R_SYMKey: ELF64_R_SYM,

		ELF64_R_TYPEKey: ELF64_R_TYPE,

		ELF64_ST_BINDKey: ELF64_ST_BIND,

		ELF64_ST_INFOKey: ELF64_ST_INFO,

		ELF64_ST_TYPEKey: ELF64_ST_TYPE,

		ELF64_ST_VISIBILITYKey: ELF64_ST_VISIBILITY,

		ElfemitrelocKey: Elfemitreloc,

		ElfinitKey: Elfinit,

		ElfwritedynentKey: Elfwritedynent,

		ElfwritedynentsymplusKey: Elfwritedynentsymplus,

		EntryvalueKey: Entryvalue,

		ErrorfKey: Errorf,

		ExitKey: Exit,

		ExitfKey: Exitf,

		LflagKey: Lflag,

		LoaderblkKey: Loaderblk,

		MachoemitrelocKey: Machoemitreloc,

		MainKey: Main,

		PeinitKey: Peinit,

		RndKey: Rnd,

		SymaddrKey: Symaddr,

		XcoffadddynrelKey: Xcoffadddynrel,

		XcoffinitKey: Xcoffinit,
	}
}

var Adddynsym = ld.Adddynsym

var Addstring = ld.Addstring

var Asmbelf = ld.Asmbelf

var Asmbelfsetup = ld.Asmbelfsetup

var Asmbmacho = ld.Asmbmacho

var Asmbpe = ld.Asmbpe

// Generate XCOFF assembly file
var Asmbxcoff = ld.Asmbxcoff

var Asmelfsym = ld.Asmelfsym

var Asmplan9sym = ld.Asmplan9sym

var AtExit = ld.AtExit

var Codeblk = ld.Codeblk

var CodeblkPad = ld.CodeblkPad

// TODO(josharian): delete. See issue 19865.
var Cputime = ld.Cputime

var Datblk = ld.Datblk

var Domacholink = ld.Domacholink

var Dwarfblk = ld.Dwarfblk

var ELF32_R_INFO = ld.ELF32_R_INFO

var ELF32_R_SYM = ld.ELF32_R_SYM

var ELF32_R_TYPE = ld.ELF32_R_TYPE

var ELF32_ST_BIND = ld.ELF32_ST_BIND

var ELF32_ST_INFO = ld.ELF32_ST_INFO

var ELF32_ST_TYPE = ld.ELF32_ST_TYPE

var ELF32_ST_VISIBILITY = ld.ELF32_ST_VISIBILITY

var ELF64_R_INFO = ld.ELF64_R_INFO

var ELF64_R_SYM = ld.ELF64_R_SYM

var ELF64_R_TYPE = ld.ELF64_R_TYPE

var ELF64_ST_BIND = ld.ELF64_ST_BIND

var ELF64_ST_INFO = ld.ELF64_ST_INFO

var ELF64_ST_TYPE = ld.ELF64_ST_TYPE

var ELF64_ST_VISIBILITY = ld.ELF64_ST_VISIBILITY

var Elfemitreloc = ld.Elfemitreloc

//  Initialize the global variable that describes the ELF header. It will be updated as
//  we write section and prog headers.
var Elfinit = ld.Elfinit

var Elfwritedynent = ld.Elfwritedynent

var Elfwritedynentsymplus = ld.Elfwritedynentsymplus

var Entryvalue = ld.Entryvalue

// Errorf logs an error message.
//
// If more than 20 errors have been printed, exit with an error.
//
// Logging an error means that on exit cmd/link will delete any
// output file and return a non-zero error code.
var Errorf = ld.Errorf

// Exit exits with code after executing all atExitFuncs.
var Exit = ld.Exit

// Exitf logs an error message then calls Exit(2).
var Exitf = ld.Exitf

var Lflag = ld.Lflag

// Create loader section and returns its size
var Loaderblk = ld.Loaderblk

var Machoemitreloc = ld.Machoemitreloc

// Main is the main entry point for the linker code.
var Main = ld.Main

var Peinit = ld.Peinit

var Rnd = ld.Rnd

var Symaddr = ld.Symaddr

// Xcoffadddynrel adds a dynamic relocation in a XCOFF file.
// This relocation will be made by the loader.
var Xcoffadddynrel = ld.Xcoffadddynrel

// Xcoffinit initialised some internal value and setups
// already known header information
var Xcoffinit = ld.Xcoffinit
