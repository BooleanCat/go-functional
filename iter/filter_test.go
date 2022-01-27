package iter_test

import (
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
)

func TestFilter(t *testing.T) {
	isEven := func(a int) bool { return a%2 == 0 }
	evens := iter.Filter[int](iter.Count(), isEven)
	assert.Equal(t, evens.Next().Unwrap(), 0)
	assert.Equal(t, evens.Next().Unwrap(), 2)
}

func TestFilterEmpty(t *testing.T) {
	isEven := func(a int) bool { return a%2 == 0 }
	evens := iter.Filter[int](iter.Exhausted[int](), isEven)
	assert.True(t, evens.Next().IsNone())
}

func TestExclude(t *testing.T) {
	isEven := func(a int) bool { return a%2 == 0 }
	evens := iter.Exclude[int](iter.Count(), isEven)
	assert.Equal(t, evens.Next().Unwrap(), 1)
	assert.Equal(t, evens.Next().Unwrap(), 3)
}
