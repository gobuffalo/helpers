package graphh

import (
	"cmd/vendor/github.com/google/pprof/internal/graph"
	"github.com/gobuffalo/helpers/hctx"
)

const (
	ComposeDotKey = "ComposeDot"

	CreateNodesKey = "CreateNodes"

	NewKey = "New"

	ShortenFunctionNameKey = "ShortenFunctionName"

	SortTagsKey = "SortTags"
)

func New() hctx.Map {
	return hctx.Map{

		ComposeDotKey: ComposeDot,

		CreateNodesKey: CreateNodes,

		NewKey: New,

		ShortenFunctionNameKey: ShortenFunctionName,

		SortTagsKey: SortTags,
	}
}

// ComposeDot creates and writes a in the DOT format to the writer, using
// the configurations given.
var ComposeDot = graph.ComposeDot

// CreateNodes creates graph nodes for all locations in a profile. It
// returns set of all nodes, plus a mapping of each location to the
// set of corresponding nodes (one per location.Line).
var CreateNodes = graph.CreateNodes

// New summarizes performance data from a profile into a graph.
var New = graph.New

// ShortenFunctionName returns a shortened version of a function&#39;s name.
var ShortenFunctionName = graph.ShortenFunctionName

// SortTags sorts a slice of tags based on their weight.
var SortTags = graph.SortTags
