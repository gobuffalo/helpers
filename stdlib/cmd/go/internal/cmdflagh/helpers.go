package cmdflagh

import (
	"cmd/go/internal/cmdflag"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AddKnownFlagsKey = "AddKnownFlags"

	FindGOFLAGSKey = "FindGOFLAGS"

	IsBoolKey = "IsBool"

	ParseKey = "Parse"

	SetBoolKey = "SetBool"

	SetIntKey = "SetInt"

	SyntaxErrorKey = "SyntaxError"
)

func New() hctx.Map {
	return hctx.Map{

		AddKnownFlagsKey: AddKnownFlags,

		FindGOFLAGSKey: FindGOFLAGS,

		IsBoolKey: IsBool,

		ParseKey: Parse,

		SetBoolKey: SetBool,

		SetIntKey: SetInt,

		SyntaxErrorKey: SyntaxError,
	}
}

// AddKnownFlags registers the flags in defns with base.AddKnownFlag.
var AddKnownFlags = cmdflag.AddKnownFlags

// FindGOFLAGS extracts and returns the flags matching defns from GOFLAGS.
// Ideally the caller would mention that the flags were from GOFLAGS
// when reporting errors, but that&#39;s too hard for now.
var FindGOFLAGS = cmdflag.FindGOFLAGS

// IsBool reports whether v is a bool flag.
var IsBool = cmdflag.IsBool

// Parse sees if argument i is present in the definitions and if so,
// returns its definition, value, and whether it consumed an extra word.
// If the flag begins (cmd.Name()+&#34;.&#34;) it is ignored for the purpose of this function.
var Parse = cmdflag.Parse

// SetBool sets the addressed boolean to the value.
var SetBool = cmdflag.SetBool

// SetInt sets the addressed integer to the value.
var SetInt = cmdflag.SetInt

// SyntaxError reports an argument syntax error and exits the program.
var SyntaxError = cmdflag.SyntaxError
