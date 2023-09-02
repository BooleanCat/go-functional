package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/internal/fakes"
	"github.com/BooleanCat/go-functional/iter"
	"github.com/BooleanCat/go-functional/option"
)

func ExampleEnumerate() {
	iterator := iter.Enumerate[string](iter.Lift([]string{"Hello", "Friend"}))
	fmt.Println(iterator.Next())
	fmt.Println(iterator.Next())
	fmt.Println(iterator.Next())

	// Output:
	// Some((0, Hello))
	// Some((1, Friend))
	// None
}

func ExampleEnumerate_method() {
	iterator := iter.Lift([]string{"Hello", "Friend"}).Enumerate()
	fmt.Println(iterator.Next())
	fmt.Println(iterator.Next())
	fmt.Println(iterator.Next())

	// Output:
	// Some((0, Hello))
	// Some((1, Friend))
	// None
}

func TestEnumerate(t *testing.T) {
	iterator := iter.Enumerate[string](iter.Lift([]string{"Hello", "Friend"}))

	assert.Equal(t, iterator.Next(), option.Some(iter.Pair[uint, string]{0, "Hello"}))
	assert.Equal(t, iterator.Next(), option.Some(iter.Pair[uint, string]{1, "Friend"}))
	assert.True(t, iterator.Next().IsNone())
}

func TestEnumerateEmpty(t *testing.T) {
	iterator := iter.Enumerate[int](iter.Exhausted[int]())
	assert.True(t, iterator.Next().IsNone())
}

func TestEnumerateExhausted(t *testing.T) {
	delegate := new(fakes.Iterator[int])
	iterator := iter.Enumerate[int](delegate)

	delegate.NextReturnsOnCall(0, option.Some(42))
	delegate.NextReturnsOnCall(1, option.None[int]())

	assert.True(t, iterator.Next().IsSome())
	assert.True(t, iterator.Next().IsNone())
	assert.True(t, iterator.Next().IsNone())
	assert.Equal(t, delegate.NextCallCount(), 2)
}
