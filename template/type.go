package template

type (
	T        interface{}
	tSlice   []interface{}
	mapFunc  func(interface{}) interface{}
	foldFunc func(interface{}, interface{}) interface{}
)

func Collect(iter Iter) []interface{} {
	return collect(iter)
}

func (f *Functor) Collect() []interface{} {
	return Collect(f.iter)
}

func Fold(iter Iter, initial interface{}, op foldFunc) interface{} {
	return fold(iter, initial, op)
}

func (f *Functor) Fold(initial T, op foldFunc) interface{} {
	return Fold(f.iter, initial, op)
}

func fromT(value T) interface{} {
	return interface{}(value)
}
