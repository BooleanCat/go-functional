// Package functional implements functional operations for slices of go primitives
package functional

// BoolSliceFunctor implements functional operations for type []bool
type BoolSliceFunctor struct {
	slice []bool
}

// LiftBoolSlice creates an BoolSliceFunctor from a []bool
func LiftBoolSlice(slice []bool) BoolSliceFunctor {
	return BoolSliceFunctor{slice: slice}
}

// Collect returns the underlying []bool
func (f BoolSliceFunctor) Collect() []bool {
	return f.slice
}

// Map returns a new BoolSliceFunctor whose underlying slice is the result of
// applying the input operation to each of its members.
func (f BoolSliceFunctor) Map(op func(bool) bool) BoolSliceFunctor {
	mapped := make([]bool, 0, len(f.slice))
	for _, i := range f.slice {
		mapped = append(mapped, op(i))
	}
	return LiftBoolSlice(mapped)
}

// Filter returns a new BoolSliceFunctor whose underlying slice has had members
// exluded that do not satisfy the input filter.
func (f BoolSliceFunctor) Filter(op func(bool) bool) BoolSliceFunctor {
	var filtered []bool
	for _, i := range f.slice {
		if op(i) {
			filtered = append(filtered, i)
		}
	}
	return LiftBoolSlice(filtered)
}

// Exclude returns a new BoolSliceFunctor whose underlying slice has had members
// exluded that satisfy the input filter.
func (f BoolSliceFunctor) Exclude(op func(bool) bool) BoolSliceFunctor {
	return LiftBoolSlice(f.slice).Filter(negateBoolOp(op))
}

// Fold applies its input operation to the initial input value and the first
// member of the underlying slice. It successively applies the input operation
// to the result of the previous and the next value in the underlying slice. It
// returns the final value successful operations. If the underlying slice is
// empty then Fold returns the initial input value.
func (f BoolSliceFunctor) Fold(initial bool, op func(bool, bool) bool) bool {
	for _, i := range f.slice {
		initial = op(initial, i)
	}
	return initial
}

// Take returns a new BoolSliceFunctor whose underlying slice has had all
// members after the nth dropped. If n is larger than the length of the
// underlying slice, Take is a no-op.
func (f BoolSliceFunctor) Take(n int) BoolSliceFunctor {
	if n > len(f.slice) {
		return f
	}
	return LiftBoolSlice(f.slice[0:n])
}

// Drop returns a new BoolSliceFunctor whose underlying slice has had the first
// n members dropped. If n is larger than the length of the underlying slice,
// Drop returns an empty StringSliceFunctor.
func (f BoolSliceFunctor) Drop(n int) BoolSliceFunctor {
	if n > len(f.slice) {
		return LiftBoolSlice([]bool{})
	}
	return LiftBoolSlice(f.slice[n:len(f.slice)])
}

// WithErrs creates an BoolSliceErrFunctor from a BoolSliceFunctor.
func (f BoolSliceFunctor) WithErrs() BoolSliceErrFunctor {
	return BoolSliceErrFunctor{slice: f.slice}
}

// BoolSliceErrFunctor behaves like BoolSliceFunctor except that operations
// performed over the underlying slice are allowed to return errors. Should
// an error occur then the BoolSliceErrFunctor's future operations do nothing
// except that Collect will return the error that occurred.
type BoolSliceErrFunctor struct {
	slice []bool
	err   error
}

// Collect returns the underlying []bool or an error if one has occurred.
func (f BoolSliceErrFunctor) Collect() ([]bool, error) {
	if f.err != nil {
		return nil, f.err
	}
	return f.slice, nil
}

// Map returns a new BoolSliceErrFunctor whose underlying slice is the result of
// applying the input operation to each of its members. Should an error occur,
// the underlying slice is lost and subsequent Collect calls will always return
// the error.
func (f BoolSliceErrFunctor) Map(op func(bool) (bool, error)) BoolSliceErrFunctor {
	if f.err != nil {
		return f
	}

	mapped := make([]bool, len(f.slice))
	for i := range f.slice {
		new, err := op(f.slice[i])
		if err != nil {
			return BoolSliceErrFunctor{err: err}
		}
		mapped[i] = new
	}
	return LiftBoolSlice(mapped).WithErrs()
}

// Filter returns a new BoolSliceErrFunctor whose underlying slice has had
// members exluded that do not satisfy the input filter. Should an error occur,
// the underlying slice is lost and subsequent Collect calls will always return
// the error.
func (f BoolSliceErrFunctor) Filter(op func(bool) (bool, error)) BoolSliceErrFunctor {
	if f.err != nil {
		return f
	}

	var filtered []bool
	for i := range f.slice {
		include, err := op(f.slice[i])
		if err != nil {
			return BoolSliceErrFunctor{err: err}
		}
		if include {
			filtered = append(filtered, f.slice[i])
		}
	}
	return LiftBoolSlice(filtered).WithErrs()
}

// Exclude returns a new BoolSliceErrFunctor whose underlying slice has had
// members exluded that satisfy the input filter. Should an error occur, the
// underlying slice is lost and subsequent Collect calls will always return the
// error.
func (f BoolSliceErrFunctor) Exclude(op func(bool) (bool, error)) BoolSliceErrFunctor {
	return LiftBoolSlice(f.slice).WithErrs().Filter(negateBoolOpWithErr(op))
}

// Fold applies its input operation to the initial input value and the first
// member of the underlying slice. It successively applies the input operation
// to the result of the previous and the next value in the underlying slice. It
// returns the final value successful operations. If the underlying slice is
// empty then Fold returns the initial input value. Should an error have
// previously occurred, that error is immediately returned.
func (f BoolSliceErrFunctor) Fold(initial bool, op func(bool, bool) bool) (bool, error) {
	if f.err != nil {
		return initial, f.err
	}

	return LiftBoolSlice(f.slice).Fold(initial, op), nil
}

// Take returns a new BoolSliceErrFunctor whose underlying slice has had all
// members after the nth dropped. If n is larger than the length of the
// underlying slice, Take is a no-op.
func (f BoolSliceErrFunctor) Take(n int) BoolSliceErrFunctor {
	if f.err != nil {
		return f
	}

	return LiftBoolSlice(f.slice).Take(n).WithErrs()
}

// Drop returns a new BoolSliceErrFunctor whose underlying slice has had the
// first n members dropped. If n is larger than the length of the underlying
// slice, Drop returns an empty BoolSliceErrFunctor.
func (f BoolSliceErrFunctor) Drop(n int) BoolSliceErrFunctor {
	if f.err != nil {
		return f
	}

	return LiftBoolSlice(f.slice).Drop(n).WithErrs()
}

func negateBoolOp(op func(bool) bool) func(bool) bool {
	return func(b bool) bool {
		return !op(b)
	}
}

func negateBoolOpWithErr(op func(bool) (bool, error)) func(bool) (bool, error) {
	return func(b bool) (bool, error) {
		result, err := op(b)
		return !result, err
	}
}
