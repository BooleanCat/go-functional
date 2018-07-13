package template

type Option struct {
	Value   T
	present bool
}

func Some(value T) Option {
	return Option{Value: value, present: true}
}

func None() Option {
	return Option{}
}

func (o Option) Present() bool {
	return o.present
}
