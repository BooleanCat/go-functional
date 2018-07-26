package template

type RepeatIter struct {
	value T
}

func NewRepeat(value T) RepeatIter {
	return RepeatIter{value}
}

func (iter RepeatIter) Next() Option {
	return Some(iter.value)
}
