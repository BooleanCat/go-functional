package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
)

func ExampleFromChannel() {
	ch := make(chan int, 2)

	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()

	fmt.Println(iter.FromChannel(ch).Collect())
	// Output: [1 2]
}

func TestFromChannel(t *testing.T) {
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()

	numbers := iter.FromChannel(ch)

	assert.Equal(t, numbers.Next().Unwrap(), 1)
	assert.Equal(t, numbers.Next().Unwrap(), 2)
	assert.Equal(t, numbers.Next().Unwrap(), 3)
	assert.True(t, numbers.Next().IsNone())
}

func TestFromChannelEmpty(t *testing.T) {
	ch := make(chan int)
	close(ch)
	assert.True(t, iter.FromChannel(ch).Next().IsNone())
}

func TestFromChannelCollect(t *testing.T) {
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()

	numbers := iter.FromChannel(ch).Collect()
	assert.SliceEqual(t, numbers, []int{1, 2})
}

func TestFromChannelDrop(t *testing.T) {
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()

	numbers := iter.FromChannel(ch).Drop(1).Collect()
	assert.SliceEqual(t, numbers, []int{2, 3})
}
