package template

type ExcludeIter struct {
	iter    Iter
	exclude filterErrFunc
}

func Exclude(iter Iter, exclude filterFunc) ExcludeIter {
	return ExcludeIter{iter, asFilterErrFunc(exclude)}
}

func ExcludeErr(iter Iter, exclude filterErrFunc) ExcludeIter {
	return ExcludeIter{iter, exclude}
}

func (iter ExcludeIter) Next() OptionalResult {
	for {
		next := iter.iter.Next()
		if next.Error() != nil || !next.Value().Present() {
			return next
		}

		result, err := iter.exclude(fromT(next.Value().Value()))
		if err != nil {
			return Failure(err)
		}

		if !result {
			return next
		}
	}
}
