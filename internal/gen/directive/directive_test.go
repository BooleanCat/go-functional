package directive_test

import (
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/internal/gen/directive"
)

func TestFromString(t *testing.T) {
	line := "//gofunctional:generate *CounterIter int Drop Take Chain"
	d := directive.FromString(line).Unwrap()
	assert.Equal(t, d.Type, "*CounterIter")
	assert.Equal(t, d.YieldedType, "int")
	assert.SliceEqual(t, d.Methods, []string{"Drop", "Take", "Chain"})
}

func TestFromStringOneMethod(t *testing.T) {
	line := "//gofunctional:generate *TakeIter[T] T Drop"
	d := directive.FromString(line).Unwrap()
	assert.Equal(t, d.Type, "*TakeIter[T]")
	assert.Equal(t, d.YieldedType, "T")
	assert.SliceEqual(t, d.Methods, []string{"Drop"})
}

func TestFromInvalid(t *testing.T) {
	line := "//gofunctional:generate *TakeIter[T] T"
	err := directive.FromString(line).UnwrapErr()
	assert.Equal(t, err.Error(), `invalid directive: //gofunctional:generate *TakeIter[T] T`)
}
