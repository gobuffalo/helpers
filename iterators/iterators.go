package iterators

import "github.com/gobuffalo/helpers/hctx"

const (
	RangeKey   = "range"
	BetweenKey = "between"
	UntilKey   = "until"
	GroupByKey = "groupBy"
)

func New() hctx.Map {
	return hctx.Map{
		RangeKey:   Range,
		BetweenKey: Between,
		UntilKey:   Until,
		GroupByKey: GroupBy,
	}
}
