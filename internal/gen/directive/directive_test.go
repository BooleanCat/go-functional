package directive_test

import (
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/internal/gen/directive"
)

func TestFromString(t *testing.T) {
	testCases := []struct {
		candidate string
		directive directive.Directive
	}{
		{
			candidate: "//gofunctional:generate *CounterIter int Drop Take Chain",
			directive: directive.Directive{Type: "*CounterIter", YieldedType: "int", Methods: []string{"Drop", "Take", "Chain"}},
		},
		{
			candidate: "//gofunctional:generate *TakeIter[T] T Drop",
			directive: directive.Directive{Type: "*TakeIter[T]", YieldedType: "T", Methods: []string{"Drop"}},
		},
		{
			candidate: "//gofunctional:generate *TakeIter[T, U, V] Tuple[U, V] Drop Collect",
			directive: directive.Directive{Type: "*TakeIter[T, U, V]", YieldedType: "Tuple[U, V]", Methods: []string{"Drop", "Collect"}},
		},
		{
			candidate: "//gofunctional:generate *LinesIter []byte Drop Collect",
			directive: directive.Directive{Type: "*LinesIter", YieldedType: "[]byte", Methods: []string{"Drop", "Collect"}},
		},
	}
	for _, tc := range testCases {
		got := directive.FromString(tc.candidate).Unwrap()
		assert.Equal(t, tc.directive.Type, got.Type)
		assert.Equal(t, tc.directive.YieldedType, got.YieldedType)
		assert.SliceEqual(t, tc.directive.Methods, got.Methods)
	}
}

func TestFromInvalid(t *testing.T) {
	line := "//gofunctional:generate *TakeIter[T] T"
	err := directive.FromString(line).UnwrapErr()
	assert.Equal(t, err.Error(), `invalid directive: //gofunctional:generate *TakeIter[T] T`)
}
