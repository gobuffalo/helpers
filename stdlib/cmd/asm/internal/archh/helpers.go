package archh

import (
	"cmd/asm/internal/arch"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ARM64RegisterArrangementKey = "ARM64RegisterArrangement"

	ARM64RegisterExtensionKey = "ARM64RegisterExtension"

	ARM64RegisterListOffsetKey = "ARM64RegisterListOffset"

	ARM64SuffixKey = "ARM64Suffix"

	ARMConditionCodesKey = "ARMConditionCodes"

	ARMMRCOffsetKey = "ARMMRCOffset"

	IsARM64CMPKey = "IsARM64CMP"

	IsARM64STLXRKey = "IsARM64STLXR"

	IsARM64TBLKey = "IsARM64TBL"

	IsARMBFXKey = "IsARMBFX"

	IsARMCMPKey = "IsARMCMP"

	IsARMFloatCmpKey = "IsARMFloatCmp"

	IsARMMRCKey = "IsARMMRC"

	IsARMMULAKey = "IsARMMULA"

	IsARMSTREXKey = "IsARMSTREX"

	IsMIPSCMPKey = "IsMIPSCMP"

	IsMIPSMULKey = "IsMIPSMUL"

	IsPPC64CMPKey = "IsPPC64CMP"

	IsPPC64ISELKey = "IsPPC64ISEL"

	IsPPC64NEGKey = "IsPPC64NEG"

	IsPPC64RLDKey = "IsPPC64RLD"

	IsS390xCMPKey = "IsS390xCMP"

	IsS390xNEGKey = "IsS390xNEG"

	ParseARMConditionKey = "ParseARMCondition"

	SetKey = "Set"
)

func New() hctx.Map {
	return hctx.Map{

		ARM64RegisterArrangementKey: ARM64RegisterArrangement,

		ARM64RegisterExtensionKey: ARM64RegisterExtension,

		ARM64RegisterListOffsetKey: ARM64RegisterListOffset,

		ARM64SuffixKey: ARM64Suffix,

		ARMConditionCodesKey: ARMConditionCodes,

		ARMMRCOffsetKey: ARMMRCOffset,

		IsARM64CMPKey: IsARM64CMP,

		IsARM64STLXRKey: IsARM64STLXR,

		IsARM64TBLKey: IsARM64TBL,

		IsARMBFXKey: IsARMBFX,

		IsARMCMPKey: IsARMCMP,

		IsARMFloatCmpKey: IsARMFloatCmp,

		IsARMMRCKey: IsARMMRC,

		IsARMMULAKey: IsARMMULA,

		IsARMSTREXKey: IsARMSTREX,

		IsMIPSCMPKey: IsMIPSCMP,

		IsMIPSMULKey: IsMIPSMUL,

		IsPPC64CMPKey: IsPPC64CMP,

		IsPPC64ISELKey: IsPPC64ISEL,

		IsPPC64NEGKey: IsPPC64NEG,

		IsPPC64RLDKey: IsPPC64RLD,

		IsS390xCMPKey: IsS390xCMP,

		IsS390xNEGKey: IsS390xNEG,

		ParseARMConditionKey: ParseARMCondition,

		SetKey: Set,
	}
}

// ARM64RegisterArrangement parses an ARM64 vector register arrangement.
var ARM64RegisterArrangement = arch.ARM64RegisterArrangement

// ARM64RegisterExtension parses an ARM64 register with extension or arrangement.
var ARM64RegisterExtension = arch.ARM64RegisterExtension

// ARM64RegisterListOffset generates offset encoding according to AArch64 specification.
var ARM64RegisterListOffset = arch.ARM64RegisterListOffset

// ARM64Suffix handles the special suffix for the ARM64.
// It returns a boolean to indicate success; failure means
// cond was unrecognized.
var ARM64Suffix = arch.ARM64Suffix

// ARMConditionCodes handles the special condition code situation for the ARM.
// It returns a boolean to indicate success; failure means cond was unrecognized.
var ARMConditionCodes = arch.ARMConditionCodes

// ARMMRCOffset implements the peculiar encoding of the MRC and MCR instructions.
// The difference between MRC and MCR is represented by a bit high in the word, not
// in the usual way by the opcode itself. Asm must use AMRC for both instructions, so
// we return the opcode for MRC so that asm doesn&#39;t need to import obj/arm.
var ARMMRCOffset = arch.ARMMRCOffset

// IsARM64CMP reports whether the op (as defined by an arm.A* constant) is
// one of the comparison instructions that require special handling.
var IsARM64CMP = arch.IsARM64CMP

// IsARM64STLXR reports whether the op (as defined by an arm64.A*
// constant) is one of the STLXR-like instructions that require special
// handling.
var IsARM64STLXR = arch.IsARM64STLXR

// IsARM64TBL reports whether the op (as defined by an arm64.A*
// constant) is one of the table lookup instructions that require special
// handling.
var IsARM64TBL = arch.IsARM64TBL

// IsARMBFX reports whether the op (as defined by an arm.A* constant) is one the
// BFX-like instructions which are in the form of &#34;op $width, $LSB, (Reg,) Reg&#34;.
var IsARMBFX = arch.IsARMBFX

// IsARMCMP reports whether the op (as defined by an arm.A* constant) is
// one of the comparison instructions that require special handling.
var IsARMCMP = arch.IsARMCMP

// IsARMFloatCmp reports whether the op is a floating comparison instruction.
var IsARMFloatCmp = arch.IsARMFloatCmp

// IsARMMRC reports whether the op (as defined by an arm.A* constant) is
// MRC or MCR
var IsARMMRC = arch.IsARMMRC

// IsARMMULA reports whether the op (as defined by an arm.A* constant) is
// MULA, MULS, MMULA, MMULS, MULABB, MULAWB or MULAWT, the 4-operand instructions.
var IsARMMULA = arch.IsARMMULA

// IsARMSTREX reports whether the op (as defined by an arm.A* constant) is
// one of the STREX-like instructions that require special handling.
var IsARMSTREX = arch.IsARMSTREX

// IsMIPSCMP reports whether the op (as defined by an mips.A* constant) is
// one of the CMP instructions that require special handling.
var IsMIPSCMP = arch.IsMIPSCMP

// IsMIPSMUL reports whether the op (as defined by an mips.A* constant) is
// one of the MUL/DIV/REM instructions that require special handling.
var IsMIPSMUL = arch.IsMIPSMUL

// IsPPC64CMP reports whether the op (as defined by an ppc64.A* constant) is
// one of the CMP instructions that require special handling.
var IsPPC64CMP = arch.IsPPC64CMP

var IsPPC64ISEL = arch.IsPPC64ISEL

// IsPPC64NEG reports whether the op (as defined by an ppc64.A* constant) is
// one of the NEG-like instructions that require special handling.
var IsPPC64NEG = arch.IsPPC64NEG

// IsPPC64RLD reports whether the op (as defined by an ppc64.A* constant) is
// one of the RLD-like instructions that require special handling.
// The FMADD-like instructions behave similarly.
var IsPPC64RLD = arch.IsPPC64RLD

// IsS390xCMP reports whether the op (as defined by an s390x.A* constant) is
// one of the CMP instructions that require special handling.
var IsS390xCMP = arch.IsS390xCMP

// IsS390xNEG reports whether the op (as defined by an s390x.A* constant) is
// one of the NEG-like instructions that require special handling.
var IsS390xNEG = arch.IsS390xNEG

// ParseARMCondition parses the conditions attached to an ARM instruction.
// The input is a single string consisting of period-separated condition
// codes, such as &#34;.P.W&#34;. An initial period is ignored.
var ParseARMCondition = arch.ParseARMCondition

// Set configures the architecture specified by GOARCH and returns its representation.
// It returns nil if GOARCH is not recognized.
var Set = arch.Set
