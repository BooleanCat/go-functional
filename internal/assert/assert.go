package assert

import "testing"

func Equal[T comparable](t *testing.T, a, b T) {
	t.Helper()

	if a != b {
		t.Errorf("expected `%v` to equal `%v`", a, b)
	}
}

func True(t *testing.T, b bool) {
	t.Helper()

	if !b {
		t.Error("expected `false` to be `true`")
	}
}

func False(t *testing.T, b bool) {
	t.Helper()

	if b {
		t.Error("expected `true` to be `false`")
	}
}

func Nil(t *testing.T, v interface{}) {
	t.Helper()

	if v != nil {
		t.Errorf("expected `%v` to equal `nil`", v)
	}
}

func NotNil(t *testing.T, v interface{}) {
	t.Helper()

	if v == nil {
		t.Error("expected `nil` not to equal `nil`")
	}
}

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
