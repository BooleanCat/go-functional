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

func (iter ExcludeIter) Next() Result {
	for {
		next := iter.iter.Next()
		if next.Error() != nil {
			return next
		}

		result, err := iter.exclude(fromT(next.Value()))
		if err != nil {
			return Failed(err)
		}

		if !result {
			return next
		}
	}
}
