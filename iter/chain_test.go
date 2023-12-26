package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/internal/fakes"
	"github.com/BooleanCat/go-functional/iter"
)

func ExampleChain() {
	fmt.Println(iter.Chain[int](iter.Lift([]int{1, 2}), iter.Lift([]int{3, 4}), iter.Lift([]int{0, 9})).Collect())
	// Output: [1 2 3 4 0 9]
}

func ExampleChain_method() {
	fmt.Println(iter.Lift([]int{1, 2}).Chain(iter.Lift([]int{3, 4}), iter.Lift([]int{0, 9})).Collect())
	// Output: [1 2 3 4 0 9]
}

func ExampleChainIter_String() {
	fmt.Println(iter.Chain[int]())
	// Output: Iterator<Chain, type=int>
}

func TestChainMultiple(t *testing.T) {
	items := iter.Chain[int](iter.Lift([]int{1, 2}), iter.Lift([]int{3, 4}))
	assert.Equal(t, items.Next().Unwrap(), 1)
	assert.Equal(t, items.Next().Unwrap(), 2)
	assert.Equal(t, items.Next().Unwrap(), 3)
	assert.Equal(t, items.Next().Unwrap(), 4)
	assert.True(t, items.Next().IsNone())
}

func TestChainSingle(t *testing.T) {
	items := iter.Chain[int](iter.Lift([]int{1, 2}))
	assert.Equal(t, items.Next().Unwrap(), 1)
	assert.Equal(t, items.Next().Unwrap(), 2)
	assert.True(t, items.Next().IsNone())
}

func TestChainEmpty(t *testing.T) {
	assert.True(t, iter.Chain[int]().Next().IsNone())
}

func TestChainExhausted(t *testing.T) {
	delegate1 := new(fakes.Iterator[int])
	delegate2 := new(fakes.Iterator[int])
	iter := iter.Chain[int](delegate1, delegate2)

	assert.True(t, iter.Next().IsNone())
	assert.True(t, iter.Next().IsNone())
	assert.Equal(t, delegate1.NextCallCount(), 1)
	assert.Equal(t, delegate2.NextCallCount(), 1)
}

func TestChainString(t *testing.T) {
	assert.Equal(t, iter.Chain[int]().String(), "Iterator<Chain, type=int>")
}
