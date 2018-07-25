package template

type (
	T      interface{}
	tSlice []interface{}
)

func Lambda(f func(interface{}) interface{}) func(T) T {
	return func(a T) T {
		return T(f(interface{}(a)))
	}
}

func Λ(f func(interface{}) interface{}) func(T) T {
	return Lambda(f)
}

func TFold(f func(interface{}, interface{}) interface{}) func(T, T) T {
	return func(a, b T) T {
		return T(f(interface{}(a), interface{}(b)))
	}
}

func Π(f func(interface{}, interface{}) interface{}) func(T, T) T {
	return TFold(f)
}
