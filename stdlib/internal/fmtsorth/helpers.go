package fmtsorth

import (
	"internal/fmtsort"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	SortKey = "Sort"
)

func New() hctx.Map {
	return hctx.Map{

		SortKey: Sort,
	}
}

// Sort accepts a map and returns a SortedMap that has the same keys and
// values but in a stable sorted order according to the keys, modulo issues
// raised by unorderable key values such as NaNs.
//
// The ordering rules are more general than with Go&#39;s &lt; operator:
//
//  - when applicable, nil compares low
//  - ints, floats, and strings order by &lt;
//  - NaN compares less than non-NaN floats
//  - bool compares false before true
//  - complex compares real, then imag
//  - pointers compare by machine address
//  - channel values compare by machine address
//  - structs compare each field in turn
//  - arrays compare each element in turn.
//    Otherwise identical arrays compare by length.
//  - interface values compare first by reflect.Type describing the concrete type
//    and then by concrete value as described in the previous rules.
var Sort = fmtsort.Sort
