package template

type Option struct {
	value   T
	present bool
}

func Some(value T) Option {
	return Option{value: value, present: true}
}

func None() Option {
	return Option{}
}

func (o Option) Value() T {
	return o.value
}

func (o Option) Present() bool {
	return o.present
}
