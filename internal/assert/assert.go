package assert

import "testing"

func SliceEqual[T comparable](t *testing.T, a, b []T) {
	t.Helper()

	if len(a) != len(b) {
		t.Errorf("expected `%v` to equal `%v` but lengths differ", a, b)
		return
	}

	for i, v := range a {
		if v != b[i] {
			t.Errorf("expected `%v` to equal `%v`", a, b)
		}
	}
}

func Empty[S any, T []S | ~string](t *testing.T, items T) {
	t.Helper()

	if len(items) != 0 {
		t.Errorf("expected `%v` to be empty", items)
	}
}
