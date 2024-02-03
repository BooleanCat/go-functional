package iter

import "fmt"

// Pairs of values.
type Pair[V, W any] struct {
	One V
	Two W
}

func (p Pair[V, W]) String() string {
	one := fmt.Sprintf("%+v", p.One)
	two := fmt.Sprintf("%+v", p.Two)

	if val, ok := interface{}(p.One).(fmt.Stringer); ok {
		one = val.String()
	}

	if val, ok := interface{}(p.Two).(fmt.Stringer); ok {
		two = val.String()
	}

	return fmt.Sprintf("(%s, %s)", one, two)
}
