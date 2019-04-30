package parseh

import (
	"text/template/parse"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	IsEmptyTreeKey = "IsEmptyTree"

	NewKey = "New"

	NewIdentifierKey = "NewIdentifier"

	ParseKey = "Parse"
)

func New() hctx.Map {
	return hctx.Map{

		IsEmptyTreeKey: IsEmptyTree,

		NewKey: New,

		NewIdentifierKey: NewIdentifier,

		ParseKey: Parse,
	}
}

// IsEmptyTree reports whether this tree (node) is empty of everything but space.
var IsEmptyTree = parse.IsEmptyTree

// New allocates a new parse tree with the given name.
var New = parse.New

// NewIdentifier returns a new IdentifierNode with the given identifier name.
var NewIdentifier = parse.NewIdentifier

// Parse returns a map from template name to parse.Tree, created by parsing the
// templates described in the argument string. The top-level template will be
// given the specified name. If an error is encountered, parsing stops and an
// empty map is returned with the error.
var Parse = parse.Parse
