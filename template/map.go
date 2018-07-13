package template

type MapIter struct {
	iter Iter
	op   func(T) T
}

func NewMap(iter Iter, op func(T) T) MapIter {
	return MapIter{iter: iter, op: op}
}

func (iter MapIter) Next() Option {
	next := iter.iter.Next()
	if !next.Present() {
		return next
	}

	return Some(iter.op(next.Value))
}
