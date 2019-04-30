package bytealgh

import (
	"internal/bytealg"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	CompareKey = "Compare"

	CompareKey = "Compare"

	CountKey = "Count"

	CountKey = "Count"

	CountStringKey = "CountString"

	CountStringKey = "CountString"

	CutoverKey = "Cutover"

	CutoverKey = "Cutover"

	CutoverKey = "Cutover"

	CutoverKey = "Cutover"

	EqualKey = "Equal"

	IndexKey = "Index"

	IndexKey = "Index"

	IndexByteKey = "IndexByte"

	IndexByteKey = "IndexByte"

	IndexByteStringKey = "IndexByteString"

	IndexByteStringKey = "IndexByteString"

	IndexStringKey = "IndexString"

	IndexStringKey = "IndexString"
)

func New() hctx.Map {
	return hctx.Map{

		CompareKey: Compare,

		CompareKey: Compare,

		CountKey: Count,

		CountKey: Count,

		CountStringKey: CountString,

		CountStringKey: CountString,

		CutoverKey: Cutover,

		CutoverKey: Cutover,

		CutoverKey: Cutover,

		CutoverKey: Cutover,

		EqualKey: Equal,

		IndexKey: Index,

		IndexKey: Index,

		IndexByteKey: IndexByte,

		IndexByteKey: IndexByte,

		IndexByteStringKey: IndexByteString,

		IndexByteStringKey: IndexByteString,

		IndexStringKey: IndexString,

		IndexStringKey: IndexString,
	}
}

// go:noescape
var Compare = bytealg.Compare

var Compare = bytealg.Compare

var Count = bytealg.Count

// go:noescape
var Count = bytealg.Count

// go:noescape
var CountString = bytealg.CountString

var CountString = bytealg.CountString

// Cutover reports the number of failures of IndexByte we should tolerate
// before switching over to Index.
// n is the number of bytes processed so far.
// See the bytes.Index implementation for details.
var Cutover = bytealg.Cutover

// Cutover reports the number of failures of IndexByte we should tolerate
// before switching over to IndexShortStr.
// n is the number of bytes processed so far.
// See the bytes.Index implementation for details.
var Cutover = bytealg.Cutover

// Cutover reports the number of failures of IndexByte we should tolerate
// before switching over to Index.
// n is the number of bytes processed so far.
// See the bytes.Index implementation for details.
var Cutover = bytealg.Cutover

// Cutover reports the number of failures of IndexByte we should tolerate
// before switching over to Index.
// n is the number of bytes processed so far.
// See the bytes.Index implementation for details.
var Cutover = bytealg.Cutover

// go:noescape
var Equal = bytealg.Equal

// Index returns the index of the first instance of b in a, or -1 if b is not present in a.
// Requires 2 &lt;= len(b) &lt;= MaxLen.
var Index = bytealg.Index

// Index returns the index of the first instance of b in a, or -1 if b is not present in a.
// Requires 2 &lt;= len(b) &lt;= MaxLen.
var Index = bytealg.Index

var IndexByte = bytealg.IndexByte

// go:noescape
var IndexByte = bytealg.IndexByte

var IndexByteString = bytealg.IndexByteString

// go:noescape
var IndexByteString = bytealg.IndexByteString

// IndexString returns the index of the first instance of b in a, or -1 if b is not present in a.
// Requires 2 &lt;= len(b) &lt;= MaxLen.
var IndexString = bytealg.IndexString

// IndexString returns the index of the first instance of b in a, or -1 if b is not present in a.
// Requires 2 &lt;= len(b) &lt;= MaxLen.
var IndexString = bytealg.IndexString
