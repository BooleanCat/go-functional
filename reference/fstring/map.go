package fstring

type MapIter struct {
	iter Iter
	op   func(string) string
}

func NewMap(iter Iter, op func(string) string) *MapIter {
	return &MapIter{iter: iter, op: op}
}

func (iter *MapIter) Next() Option {
	next := iter.iter.Next()
	if !next.Present() {
		return next
	}

	return Some(iter.op(next.Value))
}
