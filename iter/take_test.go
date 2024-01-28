package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/internal/fakes"
	"github.com/BooleanCat/go-functional/iter"
)

func ExampleTake() {
	iter := iter.Take[int](iter.Count(), 2)
	fmt.Println(iter.Next())
	fmt.Println(iter.Next())
	fmt.Println(iter.Next())
	// Output:
	// Some(0)
	// Some(1)
	// None
}

func ExampleTake_method() {
	iter := iter.Count().Take(2)
	fmt.Println(iter.Next())
	fmt.Println(iter.Next())
	fmt.Println(iter.Next())
	// Output:
	// Some(0)
	// Some(1)
	// None
}

func ExampleTakeIter_String() {
	fmt.Println(iter.Take[int](iter.Count(), 2))
	// Output: Iterator<Take, type=int>
}

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

func TestTakeExhausted(t *testing.T) {
	delegate := new(fakes.Iterator[int])
	iter := iter.Take[int](delegate, 10)

	assert.True(t, iter.Next().IsNone())
	assert.True(t, iter.Next().IsNone())
	assert.Equal(t, delegate.NextCallCount(), 1)
}

func TestTakeString(t *testing.T) {
	numbers := iter.Take[int](iter.Count(), 2)
	assert.Equal(t, fmt.Sprintf("%s", numbers), "Iterator<Take, type=int>")  //nolint:gosimple
	assert.Equal(t, fmt.Sprintf("%s", *numbers), "Iterator<Take, type=int>") //nolint:gosimple
}
