package template

type FilterIter struct {
	iter   Iter
	filter func(T) bool
}

func NewFilter(iter Iter, filter func(T) bool) FilterIter {
	return FilterIter{iter: iter, filter: filter}
}

func (iter FilterIter) Next() Option {
	for {
		option := iter.iter.Next()
		if !option.Present() || iter.filter(option.Value) {
			return option
		}
	}
}
