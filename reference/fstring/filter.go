package fstring

type FilterIter struct {
	iter   Iter
	filter func(string) bool
}

func NewFilter(iter Iter, filter func(string) bool) FilterIter {
	return FilterIter{iter: iter, filter: filter}
}

func (iter FilterIter) Next() Option {
	for {
		if option := iter.iter.Next(); !option.Present() || iter.filter(option.Value) {
			return option
		}
	}
}
