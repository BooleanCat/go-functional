package template

type Iter interface {
	Next() Result
}

func collect(iter Iter) tSlice {
	slice := tSlice{}

	for {
		next := iter.Next()
		if next.Error() == ErrNoValue {
			return slice
		}

		slice = append(slice, fromT(next.Value()))
	}
}

func fold(iter Iter, initial T, op foldFunc) T {
	result := initial
	for {
		next := iter.Next()
		if next.Error() == ErrNoValue {
			return result
		}

		result = T(op(fromT(result), fromT(next.Value())))
	}
}
