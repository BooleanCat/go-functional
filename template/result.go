package template

type Result struct {
	value   T
	present bool
}

func Some(value T) Result {
	return Result{value: value, present: true}
}

func None() Result {
	return Result{}
}

func (r Result) Value() T {
	return r.value
}

func (r Result) Present() bool {
	return r.present
}
