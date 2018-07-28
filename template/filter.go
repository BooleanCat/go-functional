package template

type FilterIter struct {
	iter   Iter
	filter filterFunc
}

func Filter(iter Iter, filter filterFunc) FilterIter {
	return FilterIter{iter: iter, filter: filter}
}

func (iter FilterIter) Next() Option {
	for {
		option := iter.iter.Next()
		if !option.Present() || iter.filter(fromT(option.Value())) {
			return option
		}
	}
}
