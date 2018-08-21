package template

import "fmt"

type Iter interface {
	Next() OptionalResult
}

func collect(iter Iter) (tSlice, error) {
	slice := tSlice{}

	for {
		next := iter.Next()
		if next.Error() != nil {
			return tSlice{}, next.Error()
		}
		if !next.Value().Present() {
			return slice, nil
		}

		slice = append(slice, fromT(next.Value().Value()))
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
		if next.Error() != nil {
			var empty T
			return empty, next.Error()
		}
		if !next.Value().Present() {
			return result, nil
		}

		applied, err := op(fromT(result), fromT(next.Value().Value()))
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
