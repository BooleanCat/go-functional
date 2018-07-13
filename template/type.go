package template

type T interface{}

func TFrom(s []interface{}) []T {
	slice := make([]T, len(s))
	for i := range s {
		slice[i] = s[i]
	}
	return slice
}

func Lambda(f func(interface{}) interface{}) func(T) T {
	return func(a T) T {
		return T(f(interface{}(a)))
	}
}

func Λ(f func(interface{}) interface{}) func(T) T {
	return Lambda(f)
}
