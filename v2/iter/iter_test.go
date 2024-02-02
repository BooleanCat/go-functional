package iter_test

import "testing"

func TestNothing(t *testing.T) {
	if true != true {
		t.Errorf("This should not happen")
	}
}
