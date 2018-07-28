package template

type DropIter struct {
	iter Iter
	n    int
}

func Drop(iter Iter, n int) *DropIter {
	return &DropIter{iter: iter, n: n}
}

func (iter *DropIter) Next() Option {
	for iter.n > 0 {
		iter.n--
		iter.iter.Next()
	}

	return iter.iter.Next()
}
