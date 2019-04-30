package objh

import (
	"cmd/internal/obj"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AddrelKey = "Addrel"

	AppendpKey = "Appendp"

	Bool2intKey = "Bool2int"

	CConvKey = "CConv"

	CConvARMKey = "CConvARM"

	DconvKey = "Dconv"

	FlushplistKey = "Flushplist"

	LinknewKey = "Linknew"

	MconvKey = "Mconv"

	NewDwarfFixupTableKey = "NewDwarfFixupTable"

	NopoutKey = "Nopout"

	RLconvKey = "RLconv"

	RconvKey = "Rconv"

	RegisterOpSuffixKey = "RegisterOpSuffix"

	RegisterOpcodeKey = "RegisterOpcode"

	RegisterRegisterKey = "RegisterRegister"

	RegisterRegisterListKey = "RegisterRegisterList"

	SortSliceKey = "SortSlice"

	SortSliceKey = "SortSlice"

	WriteObjFileKey = "WriteObjFile"
)

func New() hctx.Map {
	return hctx.Map{

		AddrelKey: Addrel,

		AppendpKey: Appendp,

		Bool2intKey: Bool2int,

		CConvKey: CConv,

		CConvARMKey: CConvARM,

		DconvKey: Dconv,

		FlushplistKey: Flushplist,

		LinknewKey: Linknew,

		MconvKey: Mconv,

		NewDwarfFixupTableKey: NewDwarfFixupTable,

		NopoutKey: Nopout,

		RLconvKey: RLconv,

		RconvKey: Rconv,

		RegisterOpSuffixKey: RegisterOpSuffix,

		RegisterOpcodeKey: RegisterOpcode,

		RegisterRegisterKey: RegisterRegister,

		RegisterRegisterListKey: RegisterRegisterList,

		SortSliceKey: SortSlice,

		SortSliceKey: SortSlice,

		WriteObjFileKey: WriteObjFile,
	}
}

var Addrel = obj.Addrel

var Appendp = obj.Appendp

var Bool2int = obj.Bool2int

// CConv formats opcode suffix bits (Prog.Scond).
var CConv = obj.CConv

// CConvARM formats ARM opcode suffix bits (mostly condition codes).
var CConvARM = obj.CConvARM

var Dconv = obj.Dconv

var Flushplist = obj.Flushplist

var Linknew = obj.Linknew

var Mconv = obj.Mconv

var NewDwarfFixupTable = obj.NewDwarfFixupTable

var Nopout = obj.Nopout

var RLconv = obj.RLconv

var Rconv = obj.Rconv

// RegisterOpSuffix assigns cconv function for formatting opcode suffixes
// when compiling for GOARCH=arch.
//
// cconv is never called with 0 argument.
var RegisterOpSuffix = obj.RegisterOpSuffix

// RegisterOpcode binds a list of instruction names
// to a given instruction number range.
var RegisterOpcode = obj.RegisterOpcode

// RegisterRegister binds a pretty-printer (Rconv) for register
// numbers to a given register number range. Lo is inclusive,
// hi exclusive (valid registers are lo through hi-1).
var RegisterRegister = obj.RegisterRegister

// RegisterRegisterList binds a pretty-printer (RLconv) for register list
// numbers to a given register list number range. Lo is inclusive,
// hi exclusive (valid register list are lo through hi-1).
var RegisterRegisterList = obj.RegisterRegisterList

var SortSlice = obj.SortSlice

var SortSlice = obj.SortSlice

var WriteObjFile = obj.WriteObjFile
