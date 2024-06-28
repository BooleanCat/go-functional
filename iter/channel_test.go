package iter_test

import (
	"fmt"
	"iter"
	"slices"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	fn "github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleFromChannel() {
	items := make(chan int)

	go func() {
		defer close(items)
		items <- 1
		items <- 2
	}()

	for number := range fn.FromChannel(items) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
}

func TestFromChannelTerminateEarly(t *testing.T) {
	t.Parallel()

	channel := make(chan int, 1)
	defer close(channel)

	channel <- 1
	numbers := fn.FromChannel(channel)

	_, stop := iter.Pull(numbers)
	stop()
}

func TestFromChannelEmpty(t *testing.T) {
	t.Parallel()

	channel := make(chan int)
	close(channel)

	assert.Empty[int](t, slices.Collect(fn.FromChannel(channel)))
}

func ExampleToChannel() {
	channel := fn.ToChannel(slices.Values([]int{1, 2, 3}))

	for number := range channel {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
	// 3
}

func ExampleToChannel_method() {
	channel := fn.Iterator[int](slices.Values([]int{1, 2, 3})).ToChannel()

	for number := range channel {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
	// 3
}

func TestToChannelEmpty(t *testing.T) {
	t.Parallel()

	for range fn.ToChannel(slices.Values([]int{})) {
		t.Error("unexpected")
	}
}
