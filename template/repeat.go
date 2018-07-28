package template

type RepeatIter struct {
	value T
}

func Repeat(value T) RepeatIter {
	return RepeatIter{value}
}

func (iter RepeatIter) Next() Option {
	return Some(iter.value)
}
