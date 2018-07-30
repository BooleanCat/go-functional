package template

type ChainIter struct {
	iters []Iter
	i     int
}

func Chain(iters ...Iter) *ChainIter {
	return &ChainIter{iters: iters}
}

func (iter *ChainIter) Next() Result {
	for {
		if len(iter.iters) <= iter.i {
			return None()
		}

		next := iter.iters[iter.i].Next()
		if !next.Present() {
			iter.i++
			continue
		}

		return next
	}
}
