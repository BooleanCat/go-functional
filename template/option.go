package template

type Option struct {
	value   T
	present bool
}

func Some(t T) Option {
	return Option{t, true}
}

func None() Option {
	return Option{}
}

func (o Option) Value() T {
	return o.value
}

func (o Option) Unwrap() T {
	if !o.present {
		panic("unwrap empty option")
	}

	return o.value
}

func (o Option) Present() bool {
	return o.present
}
