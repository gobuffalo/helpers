package arm64h

import (
	"cmd/internal/obj/arm64"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ADRKey = "ADR"

	DRconvKey = "DRconv"

	FPCCMPKey = "FPCCMP"

	FPCMPKey = "FPCMP"

	FPCVTIKey = "FPCVTI"

	FPOP1SKey = "FPOP1S"

	FPOP2SKey = "FPOP2S"

	FPOP3SKey = "FPOP3S"

	LD2STRKey = "LD2STR"

	LDSTR12UKey = "LDSTR12U"

	LDSTR9SKey = "LDSTR9S"

	LDSTXKey = "LDSTX"

	MOVCONSTKey = "MOVCONST"

	OPBITKey = "OPBIT"

	OPBLRKey = "OPBLR"

	OPBccKey = "OPBcc"

	OPDP2Key = "OPDP2"

	OPDP3Key = "OPDP3"

	SYSARG4Key = "SYSARG4"

	SYSARG5Key = "SYSARG5"

	SYSHINTKey = "SYSHINT"

	SYSOPKey = "SYSOP"
)

func New() hctx.Map {
	return hctx.Map{

		ADRKey: ADR,

		DRconvKey: DRconv,

		FPCCMPKey: FPCCMP,

		FPCMPKey: FPCMP,

		FPCVTIKey: FPCVTI,

		FPOP1SKey: FPOP1S,

		FPOP2SKey: FPOP2S,

		FPOP3SKey: FPOP3S,

		LD2STRKey: LD2STR,

		LDSTR12UKey: LDSTR12U,

		LDSTR9SKey: LDSTR9S,

		LDSTXKey: LDSTX,

		MOVCONSTKey: MOVCONST,

		OPBITKey: OPBIT,

		OPBLRKey: OPBLR,

		OPBccKey: OPBcc,

		OPDP2Key: OPDP2,

		OPDP3Key: OPDP3,

		SYSARG4Key: SYSARG4,

		SYSARG5Key: SYSARG5,

		SYSHINTKey: SYSHINT,

		SYSOPKey: SYSOP,
	}
}

var ADR = arm64.ADR

var DRconv = arm64.DRconv

var FPCCMP = arm64.FPCCMP

var FPCMP = arm64.FPCMP

var FPCVTI = arm64.FPCVTI

var FPOP1S = arm64.FPOP1S

var FPOP2S = arm64.FPOP2S

var FPOP3S = arm64.FPOP3S

var LD2STR = arm64.LD2STR

var LDSTR12U = arm64.LDSTR12U

var LDSTR9S = arm64.LDSTR9S

var LDSTX = arm64.LDSTX

var MOVCONST = arm64.MOVCONST

var OPBIT = arm64.OPBIT

var OPBLR = arm64.OPBLR

var OPBcc = arm64.OPBcc

var OPDP2 = arm64.OPDP2

var OPDP3 = arm64.OPDP3

var SYSARG4 = arm64.SYSARG4

//  form offset parameter to SYS; special register number
var SYSARG5 = arm64.SYSARG5

var SYSHINT = arm64.SYSHINT

var SYSOP = arm64.SYSOP
