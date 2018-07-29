package template

type MapIter struct {
	iter Iter
	op   mapFunc
}

func Map(iter Iter, op mapFunc) MapIter {
	return MapIter{iter, op}
}

func (iter MapIter) Next() Option {
	next := iter.iter.Next()
	if !next.Present() {
		return next
	}

	return Some(T(iter.op(fromT(next.Value()))))
}
