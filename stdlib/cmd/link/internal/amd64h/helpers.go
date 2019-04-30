package amd64h

import (
	"cmd/link/internal/amd64"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AddcallKey = "Addcall"

	InitKey = "Init"

	PADDRKey = "PADDR"
)

func New() hctx.Map {
	return hctx.Map{

		AddcallKey: Addcall,

		InitKey: Init,

		PADDRKey: PADDR,
	}
}

var Addcall = amd64.Addcall

var Init = amd64.Init

var PADDR = amd64.PADDR
