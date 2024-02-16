package fifo

type pair[V, W any] struct {
	Left  V
	Right W
}

type UnzipFifo[V, W any] struct {
	slice []pair[V, W]
	left  int
	right int
}

func New[V, W any]() UnzipFifo[V, W] {
	return UnzipFifo[V, W]{slice: make([]pair[V, W], 0, 10)}
}

func (f *UnzipFifo[V, W]) Enqueue(left V, right W) {
	f.slice = append(f.slice, pair[V, W]{left, right})
}

func (f *UnzipFifo[V, W]) DequeueLeft() (V, bool) {
	if f.left >= len(f.slice) {
		var v V
		return v, false
	}

	next := f.slice[f.left]
	f.left++

	f.maybeShrink()

	return next.Left, true
}

func (f *UnzipFifo[V, W]) DequeueRight() (W, bool) {
	if f.right >= len(f.slice) {
		var w W
		return w, false
	}

	next := f.slice[f.right]
	f.right++

	f.maybeShrink()

	return next.Right, true
}

func (f *UnzipFifo[V, W]) maybeShrink() {
	if f.left > 0 && f.right > 0 {
		smaller := min(f.left, f.right)
		f.slice = append([]pair[V, W](nil), f.slice[smaller:]...)
		f.left -= smaller
		f.right -= smaller
	}
}
