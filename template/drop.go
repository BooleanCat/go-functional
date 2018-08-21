package template

type DropIter struct {
	iter Iter
	n    int
}

func Drop(iter Iter, n int) *DropIter {
	return &DropIter{iter, n}
}

func (iter *DropIter) Next() OptionalResult {
	for iter.n > 0 {
		iter.n--
		if next := iter.iter.Next(); next.Error() != nil || !next.Value().Present() {
			return next
		}
	}

	return iter.iter.Next()
}
