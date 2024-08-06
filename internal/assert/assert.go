package assert

import (
	"slices"
	"testing"
)

func True(t *testing.T, value bool) {
	t.Helper()

	if !value {
		t.Error("expected true")
	}
}

func False(t *testing.T, value bool) {
	t.Helper()

	if value {
		t.Error("expected false")
	}
}

func Equal[T comparable](t *testing.T, a, b T) {
	t.Helper()

	if a != b {
		t.Errorf("expected `%v` to equal `%v`", a, b)
	}
}

func SliceEqual[T comparable](t *testing.T, a, b []T) {
	t.Helper()

	if !slices.Equal(a, b) {
		t.Errorf("expected `%v` to equal `%v`", a, b)
	}
}

func Empty[E any, Slice ~[]E | ~string](t *testing.T, items Slice) {
	t.Helper()

	if len(items) != 0 {
		t.Errorf("expected `%v` to be empty", items)
	}
}
