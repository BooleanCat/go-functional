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
	if result.Error() == ErrNoValue {
		return nil, true, nil
	}
	return fromT(result.Value()), false, result.Error()
}

type ConcreteIter struct {
	iter GenericIter
	f    transformFunc
}

func Transform(iter GenericIter, f transformFunc) ConcreteIter {
	return ConcreteIter{iter, f}
}

func (iter ConcreteIter) Next() Result {
	result, done, err := iter.iter.Next()
	if done {
		return None()
	}
	if err != nil {
		return Failed(err)
	}

	value, err := iter.f(result)
	if err != nil {
		return Failed(err)
	}

	return Some(T(value))
}
