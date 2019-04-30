package x86h

import (
	"cmd/internal/obj/x86"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	CanUse1InsnTLSKey = "CanUse1InsnTLS"

	EncodeRegisterRangeKey = "EncodeRegisterRange"

	ParseSuffixKey = "ParseSuffix"
)

func New() hctx.Map {
	return hctx.Map{

		CanUse1InsnTLSKey: CanUse1InsnTLS,

		EncodeRegisterRangeKey: EncodeRegisterRange,

		ParseSuffixKey: ParseSuffix,
	}
}

var CanUse1InsnTLS = x86.CanUse1InsnTLS

// EncodeRegisterRange packs [reg0-reg1] list into 64-bit value that
// is intended to be stored inside obj.Addr.Offset with TYPE_REGLIST.
var EncodeRegisterRange = x86.EncodeRegisterRange

// ParseSuffix handles the special suffix for the 386/AMD64.
// Suffix bits are stored into p.Scond.
//
// Leading &#34;.&#34; in cond is ignored.
var ParseSuffix = x86.ParseSuffix
