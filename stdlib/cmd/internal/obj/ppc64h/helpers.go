package ppc64h

import (
	"cmd/internal/obj/ppc64"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AOP_IIRRKey = "AOP_IIRR"

	AOP_IRKey = "AOP_IR"

	AOP_IRRKey = "AOP_IRR"

	AOP_IRRRKey = "AOP_IRRR"

	AOP_ISELKey = "AOP_ISEL"

	AOP_RLDICKey = "AOP_RLDIC"

	AOP_RRKey = "AOP_RR"

	AOP_RRRKey = "AOP_RRR"

	AOP_RRRIKey = "AOP_RRRI"

	AOP_RRRRKey = "AOP_RRRR"

	AOP_VIRRKey = "AOP_VIRR"

	AOP_XX1Key = "AOP_XX1"

	AOP_XX2Key = "AOP_XX2"

	AOP_XX3Key = "AOP_XX3"

	AOP_XX3IKey = "AOP_XX3I"

	AOP_XX4Key = "AOP_XX4"

	AOP_Z23IKey = "AOP_Z23I"

	DRconvKey = "DRconv"

	LOP_IRRKey = "LOP_IRR"

	LOP_RRRKey = "LOP_RRR"

	OPKey = "OP"

	OPCCKey = "OPCC"

	OPVCKey = "OPVC"

	OPVCCKey = "OPVCC"

	OPVXKey = "OPVX"

	OPVXX1Key = "OPVXX1"

	OPVXX2Key = "OPVXX2"

	OPVXX3Key = "OPVXX3"

	OPVXX4Key = "OPVXX4"

	OP_BCKey = "OP_BC"

	OP_BCRKey = "OP_BCR"

	OP_BRKey = "OP_BR"

	OP_RLWKey = "OP_RLW"
)

func New() hctx.Map {
	return hctx.Map{

		AOP_IIRRKey: AOP_IIRR,

		AOP_IRKey: AOP_IR,

		AOP_IRRKey: AOP_IRR,

		AOP_IRRRKey: AOP_IRRR,

		AOP_ISELKey: AOP_ISEL,

		AOP_RLDICKey: AOP_RLDIC,

		AOP_RRKey: AOP_RR,

		AOP_RRRKey: AOP_RRR,

		AOP_RRRIKey: AOP_RRRI,

		AOP_RRRRKey: AOP_RRRR,

		AOP_VIRRKey: AOP_VIRR,

		AOP_XX1Key: AOP_XX1,

		AOP_XX2Key: AOP_XX2,

		AOP_XX3Key: AOP_XX3,

		AOP_XX3IKey: AOP_XX3I,

		AOP_XX4Key: AOP_XX4,

		AOP_Z23IKey: AOP_Z23I,

		DRconvKey: DRconv,

		LOP_IRRKey: LOP_IRR,

		LOP_RRRKey: LOP_RRR,

		OPKey: OP,

		OPCCKey: OPCC,

		OPVCKey: OPVC,

		OPVCCKey: OPVCC,

		OPVXKey: OPVX,

		OPVXX1Key: OPVXX1,

		OPVXX2Key: OPVXX2,

		OPVXX3Key: OPVXX3,

		OPVXX4Key: OPVXX4,

		OP_BCKey: OP_BC,

		OP_BCRKey: OP_BCR,

		OP_BRKey: OP_BR,

		OP_RLWKey: OP_RLW,
	}
}

//  VX-form 2-register + ST + SIX operands
var AOP_IIRR = ppc64.AOP_IIRR

//  VX-form 1-register + SIM operands
var AOP_IR = ppc64.AOP_IR

var AOP_IRR = ppc64.AOP_IRR

//  VA-form 3-register + SHB operands
var AOP_IRRR = ppc64.AOP_IRRR

var AOP_ISEL = ppc64.AOP_ISEL

var AOP_RLDIC = ppc64.AOP_RLDIC

//  VX-form 2-register operands, r/none/r
var AOP_RR = ppc64.AOP_RR

//  the order is dest, a/s, b/imm for both arithmetic and logical operations
var AOP_RRR = ppc64.AOP_RRR

//  X-form, 3-register operands + EH field
var AOP_RRRI = ppc64.AOP_RRRI

//  VA-form 4-register operands
var AOP_RRRR = ppc64.AOP_RRRR

//  VX-form 2-register + UIM operands
var AOP_VIRR = ppc64.AOP_VIRR

//  XX1-form 3-register operands, 1 VSR operand
var AOP_XX1 = ppc64.AOP_XX1

//  XX2-form 3-register operands, 2 VSR operands
var AOP_XX2 = ppc64.AOP_XX2

//  XX3-form 3 VSR operands
var AOP_XX3 = ppc64.AOP_XX3

//  XX3-form 3 VSR operands + immediate
var AOP_XX3I = ppc64.AOP_XX3I

//  XX4-form, 4 VSR operands
var AOP_XX4 = ppc64.AOP_XX4

//  Z23-form, 3-register operands + CY field
var AOP_Z23I = ppc64.AOP_Z23I

var DRconv = ppc64.DRconv

var LOP_IRR = ppc64.LOP_IRR

var LOP_RRR = ppc64.LOP_RRR

var OP = ppc64.OP

var OPCC = ppc64.OPCC

var OPVC = ppc64.OPVC

var OPVCC = ppc64.OPVCC

var OPVX = ppc64.OPVX

var OPVXX1 = ppc64.OPVXX1

var OPVXX2 = ppc64.OPVXX2

var OPVXX3 = ppc64.OPVXX3

var OPVXX4 = ppc64.OPVXX4

var OP_BC = ppc64.OP_BC

var OP_BCR = ppc64.OP_BCR

var OP_BR = ppc64.OP_BR

var OP_RLW = ppc64.OP_RLW
