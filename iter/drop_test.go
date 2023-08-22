package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/internal/fakes"
	"github.com/BooleanCat/go-functional/iter"
	"github.com/BooleanCat/go-functional/option"
)

func ExampleDrop() {
	counter := iter.Drop[int](iter.Count(), 2)
	fmt.Println(counter.Next().Unwrap())
	// Output: 2
}

func ExampleDrop_method() {
	counter := iter.Count().Drop(2)
	fmt.Println(counter.Next().Unwrap())
	// Output: 2
}

func TestDrop(t *testing.T) {
	counter := iter.Drop[int](iter.Count(), 2)
	assert.Equal(t, counter.Next().Unwrap(), 2)
}

func TestDropExhausted(t *testing.T) {
	delegate := new(fakes.Iterator[int])
	iterator := iter.Drop[int](delegate, 5)

	assert.True(t, iterator.Next().IsNone())
	assert.True(t, iterator.Next().IsNone())
	assert.Equal(t, delegate.NextCallCount(), 1)
}

func TestDropExhaustedLater(t *testing.T) {
	delegate := new(fakes.Iterator[int])
	iterator := iter.Drop[int](delegate, 1)

	delegate.NextReturnsOnCall(0, option.Some(42))
	delegate.NextReturnsOnCall(1, option.Some(43))
	delegate.NextReturnsOnCall(2, option.None[int]())

	assert.True(t, iterator.Next().IsSome())
	assert.True(t, iterator.Next().IsNone())
	assert.True(t, iterator.Next().IsNone())
	assert.Equal(t, delegate.NextCallCount(), 3)
}

func TestDropCollect(t *testing.T) {
	numbers := iter.Drop[int](iter.Lift([]int{1, 2, 3}), 2).Collect()
	assert.SliceEqual(t, numbers, []int{3})
}

func TestDropForEach(t *testing.T) {
	total := 0

	iter.Lift([]int{1, 2, 3, 4}).Drop(1).ForEach(func(number int) {
		total += number
	})

	assert.Equal(t, total, 9)
}

func TestDropFind(t *testing.T) {
	assert.Equal(t, iter.Drop[int](iter.Count(), 5).Find(func(number int) bool {
		return number%4 == 0
	}), option.Some(8))
}

func TestDropDrop(t *testing.T) {
	counter := iter.Count().Drop(2).Drop(3)
	assert.Equal(t, counter.Next().Unwrap(), 5)
}

func TestDropTake(t *testing.T) {
	numbers := iter.Count().Drop(2).Take(3).Collect()
	assert.SliceEqual(t, numbers, []int{2, 3, 4})
}
