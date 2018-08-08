package template

type Functor struct {
	iter Iter
}

func New(iter Iter) *Functor {
	return &Functor{iter}
}

type Lifted struct {
	slice tSlice
	index int
}

func Lift(slice tSlice) *Functor {
	return &Functor{iter: &Lifted{slice: slice}}
}

func (f *Lifted) Next() Result {
	if f.index >= len(f.slice) {
		return None()
	}

	f.index++
	return Some(T(f.slice[f.index-1]))
}

func (f *Functor) Filter(filter filterFunc) *Functor {
	f.iter = Filter(f.iter, filter)
	return f
}

func (f *Functor) FilterErr(filter filterErrFunc) *Functor {
	f.iter = FilterErr(f.iter, filter)
	return f
}

func (f *Functor) Exclude(exclude filterFunc) *Functor {
	f.iter = Exclude(f.iter, exclude)
	return f
}

func (f *Functor) ExcludeErr(exclude filterErrFunc) *Functor {
	f.iter = ExcludeErr(f.iter, exclude)
	return f
}

func (f *Functor) Drop(n int) *Functor {
	f.iter = Drop(f.iter, n)
	return f
}

func (f *Functor) Take(n int) *Functor {
	f.iter = Take(f.iter, n)
	return f
}

func (f *Functor) Map(op mapFunc) *Functor {
	f.iter = Map(f.iter, op)
	return f
}

func (f *Functor) MapErr(op mapErrFunc) *Functor {
	f.iter = MapErr(f.iter, op)
	return f
}

func (f *Functor) Chain(iters ...Iter) *Functor {
	f.iter = Chain(append([]Iter{f.iter}, iters...)...)
	return f
}

func (f *Functor) Blur() TransformIter {
	return Blur(f.iter)
}
