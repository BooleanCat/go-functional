package template

type (
	T       interface{}
	tSlice  []interface{}
	mapFunc func(interface{}) interface{}
)

func fromT(value T) interface{} {
	return interface{}(value)
}

func TFold(f func(interface{}, interface{}) interface{}) func(T, T) T {
	return func(a, b T) T {
		return T(f(interface{}(a), interface{}(b)))
	}
}

func Π(f func(interface{}, interface{}) interface{}) func(T, T) T {
	return TFold(f)
}
