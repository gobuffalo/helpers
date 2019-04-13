package iterators

func Range(a, b int) Iterator {
	return &ranger{pos: a - 1, end: b}
}

type ranger struct {
	pos int
	end int
}

func (r *ranger) Next() interface{} {
	if r.pos < r.end {
		r.pos++
		return r.pos
	}
	return nil
}
