package template

type GenericIter interface {
	Next() (interface{}, error)
}

type TransformIter struct {
	iter Iter
}

func Blur(iter Iter) TransformIter {
	return TransformIter{iter}
}

func (iter TransformIter) Next() (interface{}, error) {
	result := iter.iter.Next()
	return fromT(result.Value()), result.Error()
}

type ConcreteIter struct {
	iter GenericIter
	f    transformFunc
}

func Transform(iter GenericIter, f transformFunc) ConcreteIter {
	return ConcreteIter{iter, f}
}

func (iter ConcreteIter) Next() Result {
	result, err := iter.iter.Next()
	if err != nil {
		return Failed(err)
	}

	value, err := iter.f(result)
	if err != nil {
		return Failed(err)
	}

	return Some(T(value))
}
