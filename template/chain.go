package template

type ChainIter struct {
	iters []Iter
	i     int
}

func NewChain(iters ...Iter) *ChainIter {
	return &ChainIter{iters: iters}
}

func (iter *ChainIter) Next() Option {
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
