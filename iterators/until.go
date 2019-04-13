package iterators

func Until(a int) Iterator {
	return &ranger{pos: -1, end: a - 1}
}
