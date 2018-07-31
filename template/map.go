package template

type MapIter struct {
	iter Iter
	op   mapFunc
}

func Map(iter Iter, op mapFunc) MapIter {
	return MapIter{iter, op}
}

func (iter MapIter) Next() Result {
	next := iter.iter.Next()
	if next.Error() != nil {
		return next
	}

	return Some(T(iter.op(fromT(next.Value()))))
}
