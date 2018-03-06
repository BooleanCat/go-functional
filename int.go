// Package functional implements functional operations for slices of go primitives
package functional

// IntSliceFunctor implements functional operations for type []int
type IntSliceFunctor struct {
	slice []int
}

// LiftIntSlice creates an IntSliceFunctor from a []int
func LiftIntSlice(slice []int) IntSliceFunctor {
	return IntSliceFunctor{slice: slice}
}

// Collect returns the underlying []int of the IntSliceFunctor
func (f IntSliceFunctor) Collect() []int {
	return f.slice
}

// Map returns a new IntSliceFunctor who's underlying slice is the result of
// applying the input operation to each of its members.
func (f IntSliceFunctor) Map(op func(int) int) IntSliceFunctor {
	mapped := make([]int, 0, len(f.slice))
	for _, i := range f.slice {
		mapped = append(mapped, op(i))
	}
	return LiftIntSlice(mapped)
}

// Filter returns a new IntSliceFunctor who's underlying slice has had members
// exluded that do not satisfy the input filter.
func (f IntSliceFunctor) Filter(op func(int) bool) IntSliceFunctor {
	var filtered []int
	for _, i := range f.slice {
		if op(i) {
			filtered = append(filtered, i)
		}
	}
	return LiftIntSlice(filtered)
}

// Fold applies its input operation to the initial input value and the first
// member of the underlying slice. It successively applies the input operation
// to the result of the previous and the next value in the underlying slice. It
// returns the final value successful operations. If the underlying slice is
// empty then Fold returns the initial input value.
func (f IntSliceFunctor) Fold(initial int, op func(int, int) int) int {
	for _, i := range f.slice {
		initial = op(initial, i)
	}
	return initial
}

// Take returns a new IntSliceFunctor who's underlying slice has had all
// members after the nth dropped. If n is larger than the length of the
// underlying slice, Take is a no-op.
func (f IntSliceFunctor) Take(n int) IntSliceFunctor {
	if n > len(f.slice) {
		return f
	}
	return LiftIntSlice(f.slice[0:n])
}
