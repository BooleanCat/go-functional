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

func fold(iter Iter, initial T, op foldErrFunc) (T, error) {
	result := initial
	for {
		next := iter.Next()
		if next.Error() == ErrNoValue {
			return result, nil
		}
		if next.Error() != nil {
			var empty T
			return empty, next.Error()
		}

		applied, err := op(fromT(result), fromT(next.Value()))
		if err != nil {
			var empty T
			return empty, err
		}
		result = T(applied)
	}
}

func roll(iter Iter, initial T, op foldFunc) T {
	result, err := fold(iter, initial, asFoldErrFunc(op))
	if err != nil {
		panic(fmt.Sprintf("rolled an iterator and an error was encountered: %v", err))
	}
	return result
}
