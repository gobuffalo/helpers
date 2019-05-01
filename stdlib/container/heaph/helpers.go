package heaph

import (
	"container/heap"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	FixKey = "Fix"

	InitKey = "Init"

	PopKey = "Pop"

	PushKey = "Push"

	RemoveKey = "Remove"
)

func New() hctx.Map {
	return hctx.Map{

		FixKey: Fix,

		InitKey: Init,

		PopKey: Pop,

		PushKey: Push,

		RemoveKey: Remove,
	}
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(h, i) followed by a Push of the new value.
// The complexity is O(log(n)) where n = h.Len().
var Fix = heap.Fix

// A heap must be initialized before any of the heap operations
// can be used. Init is idempotent with respect to the heap invariants
// and may be called whenever the heap invariants may have been invalidated.
// Its complexity is O(n) where n = h.Len().
var Init = heap.Init

// Pop removes the minimum element (according to Less) from the heap
// and returns it. The complexity is O(log(n)) where n = h.Len().
// It is equivalent to Remove(h, 0).
var Pop = heap.Pop

// Push pushes the element x onto the heap. The complexity is
// O(log(n)) where n = h.Len().
var Push = heap.Push

// Remove removes the element at index i from the heap.
// The complexity is O(log(n)) where n = h.Len().
var Remove = heap.Remove
