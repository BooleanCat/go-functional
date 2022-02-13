package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
	"github.com/BooleanCat/go-functional/iter/filters"
)

func ExampleFilter() {
	filtered := iter.Filter[int](iter.Lift([]int{0, 1, 0, 2}), filters.IsZero[int])
	fmt.Println(filtered.Next())
	fmt.Println(filtered.Next())
	fmt.Println(filtered.Next())

	// Output:
	// Some(0)
	// Some(0)
	// None
}

func ExampleExclude() {
	filtered := iter.Exclude[int](iter.Lift([]int{0, 1, 0, 2}), filters.IsZero[int])
	fmt.Println(filtered.Next())
	fmt.Println(filtered.Next())
	fmt.Println(filtered.Next())

	// Output:
	// Some(1)
	// Some(2)
	// None
}

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
