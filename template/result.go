package template

import "fmt"

type OptionalResult struct {
	value Option
	err   error
}

func Success(o Option) OptionalResult {
	return OptionalResult{value: o, err: nil}
}

func Failure(err error) OptionalResult {
	return OptionalResult{err: err}
}

func (r OptionalResult) Value() Option {
	return r.value
}

func (r OptionalResult) Unwrap() Option {
	if r.err != nil {
		panic(fmt.Sprintf("unwrap on failed result: %v", r.err))
	}

	return r.value
}

func (r OptionalResult) Error() error {
	return r.err
}
