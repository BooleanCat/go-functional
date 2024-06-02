package iter_test

import (
	"fmt"
	it "iter"
	"testing"

	"github.com/BooleanCat/go-functional/v2/future/slices"
	"github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleFromChannel() {
	items := make(chan int)

	go func() {
		defer close(items)
		items <- 1
		items <- 2
	}()

	for number := range iter.FromChannel(items) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
}

func TestFromChannelTerminateEarly(t *testing.T) {
	t.Parallel()

	channel := make(chan int)
	close(channel)

	_, stop := it.Pull(iter.FromChannel(channel))
	stop()
}

func ExampleToChannel() {
	channel := iter.ToChannel(slices.Values([]int{1, 2, 3}))

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

	for range iter.ToChannel(slices.Values([]int{})) {
		t.Error("unexpected")
	}
}
