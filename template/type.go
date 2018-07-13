package template

type T interface{}

func TFrom(s []interface{}) []T {
	slice := make([]T, len(s))
	for i := range s {
		slice[i] = s[i]
	}
	return slice
}
