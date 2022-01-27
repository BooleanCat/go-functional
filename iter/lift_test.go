package iter_test

import (
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
)

func TestLift(t *testing.T) {
	items := iter.Lift([]int{1, 2})
	assert.Equal(t, items.Next().Unwrap(), 1)
	assert.Equal(t, items.Next().Unwrap(), 2)
	assert.True(t, items.Next().IsNone())
}

func TestLiftEmpty(t *testing.T) {
	assert.True(t, iter.Lift([]int{}).Next().IsNone())
}
