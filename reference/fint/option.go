package fint

type Option struct {
	Value   int
	present bool
}

func Some(value int) Option {
	return Option{Value: value, present: true}
}

func None() Option {
	return Option{}
}

func (o Option) Present() bool {
	return o.present
}
