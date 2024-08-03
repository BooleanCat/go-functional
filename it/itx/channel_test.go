package itx_test

import (
	"fmt"

	"github.com/BooleanCat/go-functional/v2/it/filter"
	"github.com/BooleanCat/go-functional/v2/it/itx"
)

func ExampleFromChannel() {
	items := make(chan int)

	go func() {
		defer close(items)
		items <- 1
		items <- 2
		items <- 0
	}()

	for number := range itx.FromChannel(items).Exclude(filter.IsZero) {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
}

func ExampleIterator_ToChannel() {
	channel := itx.FromSlice([]int{1, 2, 3}).ToChannel()

	for number := range channel {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
	// 3
}
