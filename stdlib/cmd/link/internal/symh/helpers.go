package symh

import (
	"cmd/link/internal/sym"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	ABIToVersionKey = "ABIToVersion"

	NewSymbolsKey = "NewSymbols"

	RelocNameKey = "RelocName"

	SortSubKey = "SortSub"

	VersionToABIKey = "VersionToABI"
)

func New() hctx.Map {
	return hctx.Map{

		ABIToVersionKey: ABIToVersion,

		NewSymbolsKey: NewSymbols,

		RelocNameKey: RelocName,

		SortSubKey: SortSub,

		VersionToABIKey: VersionToABI,
	}
}

var ABIToVersion = sym.ABIToVersion

var NewSymbols = sym.NewSymbols

var RelocName = sym.RelocName

// SortSub sorts a linked-list (by Sub) of *Symbol by Value.
// Used for sub-symbols when loading host objects (see e.g. ldelf.go).
var SortSub = sym.SortSub

var VersionToABI = sym.VersionToABI
