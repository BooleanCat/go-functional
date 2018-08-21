package template

type MapIter struct {
	iter Iter
	op   mapErrFunc
}

func Map(iter Iter, op mapFunc) MapIter {
	return MapIter{iter, asMapErrFunc(op)}
}

func MapErr(iter Iter, op mapErrFunc) MapIter {
	return MapIter{iter, op}
}

func (iter MapIter) Next() OptionalResult {
	next := iter.iter.Next()
	if next.Error() != nil || !next.Value().Present() {
		return next
	}

	result, err := iter.op(fromT(next.Value().Value()))
	if err != nil {
		return Failure(err)
	}
	return Success(Some(T(result)))
}
