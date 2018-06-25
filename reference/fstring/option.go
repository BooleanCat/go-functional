package fstring

type Option struct {
	Value   string
	present bool
}

func Some(value string) Option {
	return Option{Value: value, present: true}
}

func None() Option {
	return Option{}
}

func (o Option) Present() bool {
	return o.present
}
