package template

type FilterIter struct {
	iter   Iter
	filter filterFunc
}

func Filter(iter Iter, filter filterFunc) FilterIter {
	return FilterIter{iter, filter}
}

func (iter FilterIter) Next() Result {
	for {
		result := iter.iter.Next()
		if result.Error() == ErrNoValue || iter.filter(fromT(result.Value())) {
			return result
		}
	}
}
