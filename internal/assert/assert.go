package assert

import "testing"

func Equal[T comparable](t *testing.T, a, b T) {
	if a != b {
		t.Errorf("expected `%v` to equal `%v`", a, b)
	}
}

func True(t *testing.T, b bool) {
	if !b {
		t.Error("expected `false` to be `true`")
	}
}

func False(t *testing.T, b bool) {
	if b {
		t.Error("expected `true` to be `false`")
	}
}

func Nil(t *testing.T, v interface{}) {
	Equal(t, v, nil)
}

func NotNil(t *testing.T, v interface{}) {
	if v == nil {
		t.Error("expected `nil` not to equal `nil`")
	}
}

func SliceEqual[T comparable](t *testing.T, a, b []T) {
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

func Empty[T any](t *testing.T, items []T) {
	if len(items) != 0 {
		t.Errorf("expected `%v` to be empty", items)
	}
}
