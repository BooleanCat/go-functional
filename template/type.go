package template

type (
	T          interface{}
	tSlice     []interface{}
	mapFunc    func(interface{}) interface{}
	foldFunc   func(interface{}, interface{}) interface{}
	filterFunc func(interface{}) bool
)

func Collect(iter Iter) ([]interface{}, error) {
	return collect(iter)
}

func (f *Functor) Collect() ([]interface{}, error) {
	return Collect(f.iter)
}

func Collapse(iter Iter) []interface{} {
	return collapse(iter)
}

func (f *Functor) Collapse() []interface{} {
	return Collapse(f.iter)
}

func Fold(iter Iter, initial interface{}, op foldFunc) interface{} {
	return fold(iter, T(initial), op)
}

func (f *Functor) Fold(initial interface{}, op foldFunc) interface{} {
	return Fold(f.iter, initial, op)
}

func fromT(value T) interface{} {
	return interface{}(value)
}
