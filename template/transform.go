package template

type GenericIter interface {
	Next() (interface{}, bool, error)
}

type TransformIter struct {
	iter Iter
}

func Blur(iter Iter) TransformIter {
	return TransformIter{iter}
}

func (iter TransformIter) Next() (interface{}, bool, error) {
	result := iter.iter.Next()
	if result.Error() != nil || result.Value().Present() {
		return fromT(result.Value().Value()), false, result.Error()
	}
	return nil, true, nil
}

type ConcreteIter struct {
	iter GenericIter
	f    transformFunc
}

func Transform(iter GenericIter, f transformFunc) ConcreteIter {
	return ConcreteIter{iter, f}
}

func (iter ConcreteIter) Next() OptionalResult {
	result, done, err := iter.iter.Next()
	if done {
		return Success(None())
	}
	if err != nil {
		return Failure(err)
	}

	value, err := iter.f(result)
	if err != nil {
		return Failure(err)
	}

	return Success(Some(T(value)))
}
