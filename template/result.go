package template

import "errors"

var ErrNoValue = errors.New("no value")

type Result struct {
	value T
	err   error
}

func Some(value T) Result {
	return Result{value: value, err: nil}
}

func None() Result {
	return Result{err: ErrNoValue}
}

func Failed(err error) Result {
	return Result{err: err}
}

func (r Result) Value() T {
	return r.value
}

func (r Result) Error() error {
	return r.err
}
