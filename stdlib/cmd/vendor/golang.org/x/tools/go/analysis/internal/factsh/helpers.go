package factsh

import (
	"cmd/vendor/golang.org/x/tools/go/analysis/internal/facts"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	DecodeKey = "Decode"
)

func New() hctx.Map {
	return hctx.Map{

		DecodeKey: Decode,
	}
}

// Decode decodes all the facts relevant to the analysis of package pkg.
// The read function reads serialized fact data from an external source
// for one of of pkg&#39;s direct imports. The empty file is a valid
// encoding of an empty fact set.
//
// It is the caller&#39;s responsibility to call gob.Register on all
// necessary fact types.
var Decode = facts.Decode
