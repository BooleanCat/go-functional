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

// Collect returns the underlying []int
func (f IntSliceFunctor) Collect() []int {
	return f.slice
}

// Map returns a new IntSliceFunctor whose underlying slice is the result of
// applying the input operation to each of its members.
func (f IntSliceFunctor) Map(op func(int) int) IntSliceFunctor {
	mapped := make([]int, 0, len(f.slice))
	for _, i := range f.slice {
		mapped = append(mapped, op(i))
	}
	return LiftIntSlice(mapped)
}

// Filter returns a new IntSliceFunctor whose underlying slice has had members
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

// Take returns a new IntSliceFunctor whose underlying slice has had all
// members after the nth dropped. If n is larger than the length of the
// underlying slice, Take is a no-op.
func (f IntSliceFunctor) Take(n int) IntSliceFunctor {
	if n > len(f.slice) {
		return f
	}
	return LiftIntSlice(f.slice[0:n])
}

// Drop returns a new IntSliceFunctor whose underlying slice has had the first
// n members dropped. If n is larger than the length of the underlying slice,
// Drop returns an empty StringSliceFunctor.
func (f IntSliceFunctor) Drop(n int) IntSliceFunctor {
	if n > len(f.slice) {
		return LiftIntSlice([]int{})
	}
	return LiftIntSlice(f.slice[n:len(f.slice)])
}

// WithErrs creates an IntSliceErrFunctor from a IntSliceFunctor.
func (f IntSliceFunctor) WithErrs() IntSliceErrFunctor {
	return IntSliceErrFunctor{slice: f.slice}
}

// IntSliceErrFunctor behaves like IntSliceFunctor except that operations
// performed over the underlying slice are allowed to return errors. Should
// an error occur then the IntSliceErrFunctor's future operations do nothing
// except that Collect will return the error that occurred.
type IntSliceErrFunctor struct {
	slice []int
	err   error
}

// Collect returns the underlying []int or an error if one has occurred.
func (f IntSliceErrFunctor) Collect() ([]int, error) {
	if f.err != nil {
		return nil, f.err
	}
	return f.slice, nil
}

// Map returns a new IntSliceErrFunctor whose underlying slice is the result of
// applying the input operation to each of its members. Should an error occur,
// the underlying slice is lost and subsequent Collect calls will always return
// the error.
func (f IntSliceErrFunctor) Map(op func(int) (int, error)) IntSliceErrFunctor {
	if f.err != nil {
		return f
	}

	mapped := make([]int, len(f.slice))
	for i := range f.slice {
		new, err := op(f.slice[i])
		if err != nil {
			return IntSliceErrFunctor{err: err}
		}
		mapped[i] = new
	}
	return LiftIntSlice(mapped).WithErrs()
}

// Filter returns a new IntSliceErrFunctor whose underlying slice has had
// members exluded that do not satisfy the input filter. Should an error occur,
// the underlying slice is lost and subsequent Collect calls with always return
// the error.
func (f IntSliceErrFunctor) Filter(op func(int) (bool, error)) IntSliceErrFunctor {
	if f.err != nil {
		return f
	}

	var filtered []int
	for i := range f.slice {
		include, err := op(f.slice[i])
		if err != nil {
			return IntSliceErrFunctor{err: err}
		}
		if include {
			filtered = append(filtered, f.slice[i])
		}
	}
	return LiftIntSlice(filtered).WithErrs()
}

// Fold applies its input operation to the initial input value and the first
// member of the underlying slice. It successively applies the input operation
// to the result of the previous and the next value in the underlying slice. It
// returns the final value successful operations. If the underlying slice is
// empty then Fold returns the initial input value. Should an error have
// previously occurred, that error is immediately returned.
func (f IntSliceErrFunctor) Fold(initial int, op func(int, int) int) (int, error) {
	if f.err != nil {
		return 0, f.err
	}

	return LiftIntSlice(f.slice).Fold(initial, op), nil
}

// Take returns a new IntSliceErrFunctor whose underlying slice has had all
// members after the nth dropped. If n is larger than the length of the
// underlying slice, Take is a no-op.
func (f IntSliceErrFunctor) Take(n int) IntSliceErrFunctor {
	if f.err != nil {
		return f
	}

	return LiftIntSlice(f.slice).Take(n).WithErrs()
}

// Drop returns a new IntSliceErrFunctor whose underlying slice has had the
// first n members dropped. If n is larger than the length of the underlying
// slice, Drop returns an empty StringSliceFunctor.
func (f IntSliceErrFunctor) Drop(n int) IntSliceErrFunctor {
	if f.err != nil {
		return f
	}

	return LiftIntSlice(f.slice).Drop(n).WithErrs()
}
