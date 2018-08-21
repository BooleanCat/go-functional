package template

type ChainIter struct {
	iters []Iter
	i     int
}

func Chain(iters ...Iter) *ChainIter {
	return &ChainIter{iters: iters}
}

func (iter *ChainIter) Next() OptionalResult {
	for {
		if len(iter.iters) <= iter.i {
			return Success(None())
		}

		next := iter.iters[iter.i].Next()
		if next.Error() != nil {
			return next
		}
		if !next.Value().Present() {
			iter.i++
			continue
		}

		return next
	}
}
