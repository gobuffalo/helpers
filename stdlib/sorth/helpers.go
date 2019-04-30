package sorth

import (
	"sort"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	Float64sKey = "Float64s"

	Float64sAreSortedKey = "Float64sAreSorted"

	IntsKey = "Ints"

	IntsAreSortedKey = "IntsAreSorted"

	IsSortedKey = "IsSorted"

	ReverseKey = "Reverse"

	SearchKey = "Search"

	SearchFloat64sKey = "SearchFloat64s"

	SearchIntsKey = "SearchInts"

	SearchStringsKey = "SearchStrings"

	SliceKey = "Slice"

	SliceIsSortedKey = "SliceIsSorted"

	SliceStableKey = "SliceStable"

	SortKey = "Sort"

	StableKey = "Stable"

	StringsKey = "Strings"

	StringsAreSortedKey = "StringsAreSorted"
)

func New() hctx.Map {
	return hctx.Map{

		Float64sKey: Float64s,

		Float64sAreSortedKey: Float64sAreSorted,

		IntsKey: Ints,

		IntsAreSortedKey: IntsAreSorted,

		IsSortedKey: IsSorted,

		ReverseKey: Reverse,

		SearchKey: Search,

		SearchFloat64sKey: SearchFloat64s,

		SearchIntsKey: SearchInts,

		SearchStringsKey: SearchStrings,

		SliceKey: Slice,

		SliceIsSortedKey: SliceIsSorted,

		SliceStableKey: SliceStable,

		SortKey: Sort,

		StableKey: Stable,

		StringsKey: Strings,

		StringsAreSortedKey: StringsAreSorted,
	}
}

// Float64s sorts a slice of float64s in increasing order
// (not-a-number values are treated as less than other values).
var Float64s = sort.Float64s

// Float64sAreSorted tests whether a slice of float64s is sorted in increasing order
// (not-a-number values are treated as less than other values).
var Float64sAreSorted = sort.Float64sAreSorted

// Ints sorts a slice of ints in increasing order.
var Ints = sort.Ints

// IntsAreSorted tests whether a slice of ints is sorted in increasing order.
var IntsAreSorted = sort.IntsAreSorted

// IsSorted reports whether data is sorted.
var IsSorted = sort.IsSorted

// Reverse returns the reverse order for data.
var Reverse = sort.Reverse

// Search uses binary search to find and return the smallest index i
// in [0, n) at which f(i) is true, assuming that on the range [0, n),
// f(i) == true implies f(i+1) == true. That is, Search requires that
// f is false for some (possibly empty) prefix of the input range [0, n)
// and then true for the (possibly empty) remainder; Search returns
// the first true index. If there is no such index, Search returns n.
// (Note that the &#34;not found&#34; return value is not -1 as in, for instance,
// strings.Index.)
// Search calls f(i) only for i in the range [0, n).
//
// A common use of Search is to find the index i for a value x in
// a sorted, indexable data structure such as an array or slice.
// In this case, the argument f, typically a closure, captures the value
// to be searched for, and how the data structure is indexed and
// ordered.
//
// For instance, given a slice data sorted in ascending order,
// the call Search(len(data), func(i int) bool { return data[i] &gt;= 23 })
// returns the smallest index i such that data[i] &gt;= 23. If the caller
// wants to find whether 23 is in the slice, it must test data[i] == 23
// separately.
//
// Searching data sorted in descending order would use the &lt;=
// operator instead of the &gt;= operator.
//
// To complete the example above, the following code tries to find the value
// x in an integer slice data sorted in ascending order:
//
// 	x := 23
// 	i := sort.Search(len(data), func(i int) bool { return data[i] &gt;= x })
// 	if i &lt; len(data) &amp;&amp; data[i] == x {
// 		// x is present at data[i]
// 	} else {
// 		// x is not present in data,
// 		// but i is the index where it would be inserted.
// 	}
//
// As a more whimsical example, this program guesses your number:
//
// 	func GuessingGame() {
// 		var s string
// 		fmt.Printf(&#34;Pick an integer from 0 to 100.\n&#34;)
// 		answer := sort.Search(100, func(i int) bool {
// 			fmt.Printf(&#34;Is your number &lt;= %d? &#34;, i)
// 			fmt.Scanf(&#34;%s&#34;, &amp;s)
// 			return s != &#34;&#34; &amp;&amp; s[0] == &#39;y&#39;
// 		})
// 		fmt.Printf(&#34;Your number is %d.\n&#34;, answer)
// 	}
var Search = sort.Search

// SearchFloat64s searches for x in a sorted slice of float64s and returns the index
// as specified by Search. The return value is the index to insert x if x is not
// present (it could be len(a)).
// The slice must be sorted in ascending order.
var SearchFloat64s = sort.SearchFloat64s

// SearchInts searches for x in a sorted slice of ints and returns the index
// as specified by Search. The return value is the index to insert x if x is
// not present (it could be len(a)).
// The slice must be sorted in ascending order.
var SearchInts = sort.SearchInts

// SearchStrings searches for x in a sorted slice of strings and returns the index
// as specified by Search. The return value is the index to insert x if x is not
// present (it could be len(a)).
// The slice must be sorted in ascending order.
var SearchStrings = sort.SearchStrings

// Slice sorts the provided slice given the provided less function.
//
// The sort is not guaranteed to be stable. For a stable sort, use
// SliceStable.
//
// The function panics if the provided interface is not a slice.
var Slice = sort.Slice

// SliceIsSorted tests whether a slice is sorted.
//
// The function panics if the provided interface is not a slice.
var SliceIsSorted = sort.SliceIsSorted

// SliceStable sorts the provided slice given the provided less
// function while keeping the original order of equal elements.
//
// The function panics if the provided interface is not a slice.
var SliceStable = sort.SliceStable

// Sort sorts data.
// It makes one call to data.Len to determine n, and O(n*log(n)) calls to
// data.Less and data.Swap. The sort is not guaranteed to be stable.
var Sort = sort.Sort

// Stable sorts data while keeping the original order of equal elements.
//
// It makes one call to data.Len to determine n, O(n*log(n)) calls to
// data.Less and O(n*log(n)*log(n)) calls to data.Swap.
var Stable = sort.Stable

// Strings sorts a slice of strings in increasing order.
var Strings = sort.Strings

// StringsAreSorted tests whether a slice of strings is sorted in increasing order.
var StringsAreSorted = sort.StringsAreSorted
