package mipsh

import (
	"cmd/internal/obj/mips"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	BCONDKey = "BCOND"

	DRconvKey = "DRconv"

	FPDKey = "FPD"

	FPFKey = "FPF"

	FPVKey = "FPV"

	FPWKey = "FPW"

	MMUKey = "MMU"

	OPKey = "OP"

	OP_FRRRKey = "OP_FRRR"

	OP_IRRKey = "OP_IRR"

	OP_JMPKey = "OP_JMP"

	OP_RRRKey = "OP_RRR"

	OP_SRRKey = "OP_SRR"

	SPKey = "SP"
)

func New() hctx.Map {
	return hctx.Map{

		BCONDKey: BCOND,

		DRconvKey: DRconv,

		FPDKey: FPD,

		FPFKey: FPF,

		FPVKey: FPV,

		FPWKey: FPW,

		MMUKey: MMU,

		OPKey: OP,

		OP_FRRRKey: OP_FRRR,

		OP_IRRKey: OP_IRR,

		OP_JMPKey: OP_JMP,

		OP_RRRKey: OP_RRR,

		OP_SRRKey: OP_SRR,

		SPKey: SP,
	}
}

var BCOND = mips.BCOND

var DRconv = mips.DRconv

var FPD = mips.FPD

var FPF = mips.FPF

var FPV = mips.FPV

var FPW = mips.FPW

var MMU = mips.MMU

var OP = mips.OP

var OP_FRRR = mips.OP_FRRR

var OP_IRR = mips.OP_IRR

var OP_JMP = mips.OP_JMP

var OP_RRR = mips.OP_RRR

var OP_SRR = mips.OP_SRR

var SP = mips.SP
