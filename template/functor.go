package template

type Functor struct {
	iter Iter
}

func New(iter Iter) *Functor {
	return &Functor{iter: iter}
}

type Lifted struct {
	slice []T
	index int
}

func Lift(slice []T) *Functor {
	return &Functor{iter: &Lifted{slice: slice}}
}

func (f *Lifted) Next() Option {
	if f.index >= len(f.slice) {
		return None()
	}

	f.index++
	return Some(f.slice[f.index-1])
}

func (f *Functor) Filter(filter func(T) bool) *Functor {
	f.iter = NewFilter(f.iter, filter)
	return f
}

func (f *Functor) Exclude(exclude func(T) bool) *Functor {
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

func (f *Functor) Map(op func(T) T) *Functor {
	f.iter = NewMap(f.iter, op)
	return f
}

func (f *Functor) Chain(iter Iter) *Functor {
	f.iter = NewChain(f.iter, iter)
	return f
}

func (f *Functor) Fold(initial T, op foldOp) T {
	return Fold(f.iter, initial, op)
}

func (f *Functor) Collect() []T {
	return Collect(f.iter)
}
