package it_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/it"
)

func ExampleFromChannel() {
	items := make(chan int)

	go func() {
		defer close(items)
		items <- 1
		items <- 2
	}()

	for number := range it.FromChannel(items) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
}

func TestFromChannelYieldFalse(t *testing.T) {
	t.Parallel()

	numbersChan := make(chan int, 1)
	defer close(numbersChan)

	numbersChan <- 1
	numbers := it.FromChannel(numbersChan)

	numbers(func(value int) bool {
		return false
	})
}

func TestFromChannelEmpty(t *testing.T) {
	t.Parallel()

	channel := make(chan int)
	close(channel)

	assert.Empty[int](t, slices.Collect(it.FromChannel(channel)))
}

func ExampleToChannel() {
	channel := it.ToChannel(slices.Values([]int{1, 2, 3}))

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

	for range it.ToChannel(slices.Values([]int{})) {
		t.Error("unexpected")
	}
}
