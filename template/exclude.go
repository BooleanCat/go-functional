package template

type ExcludeIter struct {
	iter    Iter
	exclude filterFunc
}

func Exclude(iter Iter, exclude filterFunc) ExcludeIter {
	return ExcludeIter{iter, exclude}
}

func (iter ExcludeIter) Next() Result {
	for {
		if option := iter.iter.Next(); !option.Present() || !iter.exclude(fromT(option.Value())) {
			return option
		}
	}
}
