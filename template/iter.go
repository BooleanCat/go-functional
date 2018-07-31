package template

import "fmt"

type Iter interface {
	Next() Result
}

func collect(iter Iter) (tSlice, error) {
	slice := tSlice{}

	for {
		next := iter.Next()
		if next.Error() == ErrNoValue {
			return slice, nil
		}
		if next.Error() != nil {
			return tSlice{}, next.Error()
		}

		slice = append(slice, fromT(next.Value()))
	}
}

func collapse(iter Iter) tSlice {
	slice, err := collect(iter)
	if err != nil {
		panic(fmt.Sprintf("collpased an iterator and an error was encountered: %v", err))
	}
	return slice
}

func fold(iter Iter, initial T, op foldFunc) (T, error) {
	result := initial
	for {
		next := iter.Next()
		if next.Error() == ErrNoValue {
			return result, nil
		}

		result = T(op(fromT(result), fromT(next.Value())))
	}
}
