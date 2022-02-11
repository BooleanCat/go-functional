package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
)

func ExampleDrop() {
	counter := iter.Drop[int](iter.Count(), 2)
	fmt.Println(counter.Next().Unwrap())
	// Output: 2
}

func TestDrop(t *testing.T) {
	counter := iter.Drop[int](iter.Count(), 2)
	assert.Equal(t, counter.Next().Unwrap(), 2)
}

func TestDropEmpty(t *testing.T) {
	assert.True(t, iter.Drop[int](iter.Exhausted[int](), 5).Next().IsNone())
}
