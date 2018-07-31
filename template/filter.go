package template

type FilterIter struct {
	iter   Iter
	filter filterErrFunc
}

func Filter(iter Iter, filter filterFunc) FilterIter {
	return FilterIter{iter, asFilterErrFunc(filter)}
}

func FilterErr(iter Iter, filter filterErrFunc) FilterIter {
	return FilterIter{iter, filter}
}

func (iter FilterIter) Next() Result {
	for {
		next := iter.iter.Next()
		if next.Error() != nil {
			return next
		}

		result, err := iter.filter(fromT(next.Value()))
		if err != nil {
			return Failed(err)
		}

		if result {
			return next
		}
	}
}
