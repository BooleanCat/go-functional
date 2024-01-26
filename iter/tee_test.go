package iter_test

import (
	"sync"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/internal/fakes"
	"github.com/BooleanCat/go-functional/iter"
)

func TestTee(t *testing.T) {
	numbers := iter.Count().Take(3)
	one, two := numbers.Tee()

	assert.SliceEqual[int](t, one.Collect(), []int{0, 1, 2})
	assert.SliceEqual[int](t, two.Collect(), []int{0, 1, 2})
}

func TestTeeTwoTakesFirst(t *testing.T) {
	numbers := iter.Count().Take(3)
	one, two := numbers.Tee()

	assert.SliceEqual[int](t, one.Collect(), []int{0, 1, 2})
	assert.SliceEqual[int](t, two.Collect(), []int{0, 1, 2})
}

func TestTeeEmpty(t *testing.T) {
	one, two := iter.Exhausted[int]().Tee()

	assert.Empty[int](t, one.Collect())
	assert.Empty[int](t, two.Collect())
}

func TestTeeParallel(t *testing.T) {
	one, two := iter.Count().Take(100000).Tee()

	wait := sync.WaitGroup{}
	wait.Add(2)

	go func() {
		defer wait.Done()
		assert.SliceEqual[int](t, one.Collect(), iter.Count().Take(100000).Collect())
	}()

	go func() {
		defer wait.Done()
		assert.SliceEqual[int](t, two.Collect(), iter.Count().Take(100000).Collect())
	}()

	wait.Wait()
}

func TestTeeExhausted(t *testing.T) {
	delegate := new(fakes.Iterator[int])
	one, two := iter.Take[int](delegate, 10).Tee()

	assert.Empty[int](t, one.Collect())
	assert.Empty[int](t, two.Collect())
	assert.Equal(t, delegate.NextCallCount(), 1)
}
