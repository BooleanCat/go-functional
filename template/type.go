package template

type (
	T             interface{}
	tSlice        []interface{}
	mapFunc       func(interface{}) interface{}
	mapErrFunc    func(interface{}) (interface{}, error)
	foldFunc      func(interface{}, interface{}) interface{}
	foldErrFunc   func(interface{}, interface{}) (interface{}, error)
	filterFunc    func(interface{}) bool
	filterErrFunc func(interface{}) (bool, error)
)

func Collect(iter Iter) ([]interface{}, error) {
	return collect(iter)
}

func (f *Functor) Collect() ([]interface{}, error) {
	return collect(f.iter)
}

func Collapse(iter Iter) []interface{} {
	return collapse(iter)
}

func (f *Functor) Collapse() []interface{} {
	return collapse(f.iter)
}

func Fold(iter Iter, initial interface{}, op foldErrFunc) (interface{}, error) {
	result, err := fold(iter, T(initial), op)
	return fromT(result), err
}

func (f *Functor) Fold(initial interface{}, op foldErrFunc) (interface{}, error) {
	return Fold(f.iter, initial, op)
}

func Roll(iter Iter, initial interface{}, op foldFunc) interface{} {
	return fromT(roll(iter, initial, op))
}

func (f Functor) Roll(initial interface{}, op foldFunc) interface{} {
	return Roll(f.iter, initial, op)
}

func fromT(value T) interface{} {
	return interface{}(value)
}

func asMapErrFunc(f mapFunc) mapErrFunc {
	return func(v interface{}) (interface{}, error) {
		return f(v), nil
	}
}

func asFilterErrFunc(f filterFunc) filterErrFunc {
	return func(v interface{}) (bool, error) {
		return f(v), nil
	}
}

func asFoldErrFunc(f foldFunc) foldErrFunc {
	return func(v, w interface{}) (interface{}, error) {
		return f(v, w), nil
	}
}
