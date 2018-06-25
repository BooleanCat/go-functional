package fstring

type Functor struct {
	iter Iter
}

type Lifted struct {
	slice []string
	index int
}

func Lift(slice []string) *Functor {
	return &Functor{iter: &Lifted{slice: slice}}
}

func (f *Lifted) Next() Option {
	if f.index >= len(f.slice) {
		return None()
	}

	f.index++
	return Some(f.slice[f.index-1])
}

func (f *Functor) Filter(filter func(value string) bool) *Functor {
	f.iter = NewFilter(f.iter, filter)
	return f
}

func (f *Functor) Exclude(exclude func(value string) bool) *Functor {
	f.iter = NewExclude(f.iter, exclude)
	return f
}

func (f *Functor) Drop(n int) *Functor {
	f.iter = NewDrop(f.iter, n)
	return f
}

func (f *Functor) Take(n int) *Functor {
	f.iter = NewTake(f.iter, n)
	return f
}

func (f *Functor) Map(op func(string) string) *Functor {
	f.iter = NewMap(f.iter, op)
	return f
}

func (f *Functor) Fold(initial string, op foldOp) string {
	return Fold(f.iter, initial, op)
}

func (f *Functor) Collect() []string {
	return Collect(f.iter)
}
