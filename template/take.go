package template

type TakeIter struct {
	iter Iter
	n    int
}

func NewTake(iter Iter, n int) *TakeIter {
	return &TakeIter{iter: iter, n: n}
}

func (iter *TakeIter) Next() Option {
	if iter.n <= 0 {
		return None()
	}

	iter.n--
	return iter.iter.Next()
}
