package template

type ExcludeIter struct {
	iter    Iter
	exclude filterFunc
}

func Exclude(iter Iter, exclude filterFunc) ExcludeIter {
	return ExcludeIter{iter: iter, exclude: exclude}
}

func (iter ExcludeIter) Next() Option {
	for {
		if option := iter.iter.Next(); !option.Present() || !iter.exclude(fromT(option.Value())) {
			return option
		}
	}
}
