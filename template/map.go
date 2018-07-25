package template

type MapIter struct {
	iter Iter
	op   mapFunc
}

func NewMap(iter Iter, op mapFunc) MapIter {
	return MapIter{iter: iter, op: op}
}

func (iter MapIter) Next() Option {
	next := iter.iter.Next()
	if !next.Present() {
		return next
	}

	return Some(T(iter.op(fromT(next.Value))))
}
