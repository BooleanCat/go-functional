package iter_test

import (
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
)

func TestCount(t *testing.T) {
	counter := iter.Count()
	assert.Equal(t, counter.Next().Unwrap(), 0)
	assert.Equal(t, counter.Next().Unwrap(), 1)
	assert.Equal(t, counter.Next().Unwrap(), 2)
}
