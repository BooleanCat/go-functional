package template

type Iter interface {
	Next() Result
}

func collect(iter Iter) tSlice {
	slice := tSlice{}

	for {
		option := iter.Next()
		if !option.Present() {
			return slice
		}

		slice = append(slice, fromT(option.Value()))
	}
}

func fold(iter Iter, initial T, op foldFunc) T {
	result := initial
	for {
		next := iter.Next()
		if !next.Present() {
			return result
		}

		result = T(op(fromT(result), fromT(next.Value())))
	}
}
