package iter_test

import (
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
)

func TestTakeIter(t *testing.T) {
	iter := iter.Take[int](iter.Count(), 2)
	assert.Equal(t, iter.Next().Unwrap(), 0)
	assert.Equal(t, iter.Next().Unwrap(), 1)
	assert.True(t, iter.Next().IsNone())
}

func TestTakeIterEmpty(t *testing.T) {
	iter := iter.Take[int](iter.Count(), 0)
	assert.True(t, iter.Next().IsNone())
}
