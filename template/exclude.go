package template

type ExcludeIter struct {
	iter    Iter
	exclude func(T) bool
}

func NewExclude(iter Iter, exclude func(T) bool) ExcludeIter {
	return ExcludeIter{iter: iter, exclude: exclude}
}

func (iter ExcludeIter) Next() Option {
	for {
		if option := iter.iter.Next(); !option.Present() || !iter.exclude(option.Value) {
			return option
		}
	}
}
