package iterators

func Between(a, b int) Iterator {
	return &ranger{pos: a, end: b - 1}
}
