package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
)

func ExampleCount() {
	counter := iter.Count()
	fmt.Println(counter.Next())
	fmt.Println(counter.Next())
	fmt.Println(counter.Next())

	// Output:
	// Some(0)
	// Some(1)
	// Some(2)
}

func TestCount(t *testing.T) {
	counter := iter.Count()
	assert.Equal(t, counter.Next().Unwrap(), 0)
	assert.Equal(t, counter.Next().Unwrap(), 1)
	assert.Equal(t, counter.Next().Unwrap(), 2)
}

func TestCountDrop(t *testing.T) {
	counter := iter.Count().Drop(5)
	assert.Equal(t, counter.Next().Unwrap(), 5)
}

func TestCountTake(t *testing.T) {
	numbers := iter.Count().Take(3).Collect()
	assert.SliceEqual(t, numbers, []int{0, 1, 2})
}
