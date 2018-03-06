// Package functional implements functional operations for slices of go primitives
package functional

// StringSliceFunctor implements functional operations for type []string
type StringSliceFunctor struct {
	slice []string
}

// LiftStringSlice creates an StringSliceFunctor from a []string
func LiftStringSlice(slice []string) StringSliceFunctor {
	return StringSliceFunctor{slice: slice}
}

// Collect returns the underlying []string of the StringSliceFunctor
func (f StringSliceFunctor) Collect() []string {
	return f.slice
}

// Map returns a new StringSliceFunctor who's underlying slice is the result of
// applying the input operation to each of its members.
func (f StringSliceFunctor) Map(op func(string) string) StringSliceFunctor {
	mapped := make([]string, 0, len(f.slice))
	for _, i := range f.slice {
		mapped = append(mapped, op(i))
	}
	return LiftStringSlice(mapped)
}

// Filter returns a new StringSliceFunctor who's underlying slice has had
// members exluded that do not satisfy the input filter.
func (f StringSliceFunctor) Filter(op func(string) bool) StringSliceFunctor {
	var filtered []string
	for _, i := range f.slice {
		if op(i) {
			filtered = append(filtered, i)
		}
	}
	return LiftStringSlice(filtered)
}

// Fold applies its input operation to the initial input value and the first
// member of the underlying slice. It successively applies the input operation
// to the result of the previous and the next value in the underlying slice. It
// returns the final value successful operations. If the underlying slice is
// empty then Fold returns the initial input value.
func (f StringSliceFunctor) Fold(initial string, op func(string, string) string) string {
	for _, i := range f.slice {
		initial = op(initial, i)
	}
	return initial
}
