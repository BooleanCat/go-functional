package fstring

type Iter interface {
	Next() Option
}

func Collect(iter Iter) []string {
	slice := []string{}
	for {
		option := iter.Next()
		if !option.Present() {
			return slice
		}

		slice = append(slice, option.Value)
	}
}

func Fold(iter Iter, initial string, op foldOp) string {
	result := initial
	for {
		next := iter.Next()
		if !next.Present() {
			return result
		}

		result = op(result, next.Value)
	}
}

type foldOp func(string, string) string
