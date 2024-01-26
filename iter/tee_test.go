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
	iters := numbers.Tee()

	assert.SliceEqual[int](t, iters.One.Collect(), []int{0, 1, 2})
	assert.SliceEqual[int](t, iters.Two.Collect(), []int{0, 1, 2})
}

func TestTeeTwoTakesFirst(t *testing.T) {
	numbers := iter.Count().Take(3)
	iters := numbers.Tee()

	assert.SliceEqual[int](t, iters.Two.Collect(), []int{0, 1, 2})
	assert.SliceEqual[int](t, iters.One.Collect(), []int{0, 1, 2})
}

func TestTeeEmpty(t *testing.T) {
	iters := iter.Exhausted[int]().Tee()

	assert.Empty[int](t, iters.One.Collect())
	assert.Empty[int](t, iters.Two.Collect())
}

func TestTeeParallel(t *testing.T) {
	iters := iter.Count().Take(100000).Tee()

	wait := sync.WaitGroup{}
	wait.Add(2)

	go func() {
		defer wait.Done()
		assert.SliceEqual[int](t, iters.One.Collect(), iter.Count().Take(100000).Collect())
	}()

	go func() {
		defer wait.Done()
		assert.SliceEqual[int](t, iters.Two.Collect(), iter.Count().Take(100000).Collect())
	}()

	wait.Wait()
}

func TestTeeExhausted(t *testing.T) {
	delegate := new(fakes.Iterator[int])
	iters := iter.Take[int](delegate, 10).Tee()

	assert.Empty[int](t, iters.One.Collect())
	assert.Empty[int](t, iters.Two.Collect())
	assert.Equal(t, delegate.NextCallCount(), 1)
}
