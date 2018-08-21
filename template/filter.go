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

func (iter FilterIter) Next() OptionalResult {
	for {
		next := iter.iter.Next()
		if next.Error() != nil || !next.Value().Present() {
			return next
		}

		result, err := iter.filter(fromT(next.Value().Value()))
		if err != nil {
			return Failure(err)
		}

		if result {
			return next
		}
	}
}
