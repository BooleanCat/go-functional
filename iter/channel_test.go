package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
	"github.com/BooleanCat/go-functional/option"
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

func TestFromChannelFind(t *testing.T) {
	ch := make(chan int)

	go func() {
		defer close(ch)
		ch <- 1
		ch <- 2
		ch <- 3
	}()

	numbers := iter.FromChannel(ch)
	defer numbers.Collect()

	number := numbers.Find(func(number int) bool {
		return number == 2
	})

	assert.Equal(t, number, option.Some(2))
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

func TestFromChannelTake(t *testing.T) {
	ch := make(chan int)

	go func() {
		defer close(ch)
		ch <- 1
		ch <- 2
		ch <- 3
	}()

	iter := iter.FromChannel(ch)
	numbers := iter.Take(2).Collect()
	assert.SliceEqual(t, numbers, []int{1, 2})
	iter.Collect() // To close the channel
}
