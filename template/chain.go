package template

type ChainIter struct {
	first  Iter
	second Iter
}

func NewChain(first, second Iter) *ChainIter {
	return &ChainIter{first, second}
}

func (iter *ChainIter) Next() Option {
	next := iter.first.Next()
	if next.Present() {
		return next
	}

	return iter.second.Next()
}
