package template

type TakeIter struct {
	iter Iter
	n    int
}

func Take(iter Iter, n int) *TakeIter {
	return &TakeIter{iter, n}
}

func (iter *TakeIter) Next() Result {
	if iter.n <= 0 {
		return None()
	}

	iter.n--
	return iter.iter.Next()
}
