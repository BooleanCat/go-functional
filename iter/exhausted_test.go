package iter_test

import (
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
)

func TestExhausted(t *testing.T) {
	assert.True(t, iter.Exhausted[int]().Next().IsNone())
}
