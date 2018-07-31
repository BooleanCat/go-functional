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

func (iter MapIter) Next() Result {
	next := iter.iter.Next()
	if next.Error() != nil {
		return next
	}

	result, err := iter.op(fromT(next.Value()))
	if err != nil {
		return Failed(err)
	}
	return Some(T(result))
}
