package template

type Iter interface {
	Next() Option
}

func collect(iter Iter) tSlice {
	slice := tSlice{}

	for {
		option := iter.Next()
		if !option.Present() {
			return slice
		}

		slice = append(slice, fromT(option.Value))
	}
}

func Fold(iter Iter, initial T, op foldOp) T {
	result := initial
	for {
		next := iter.Next()
		if !next.Present() {
			return result
		}

		result = op(result, next.Value)
	}
}

type foldOp func(T, T) T
