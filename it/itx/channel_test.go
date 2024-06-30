package itx_test

import (
	"fmt"

	"github.com/BooleanCat/go-functional/v2/it/itx"
)

func ExampleIterator_ToChannel() {
	channel := itx.SlicesValues([]int{1, 2, 3}).ToChannel()

	for number := range channel {
		fmt.Println(number)
	}

	// Output:
	// 1
	// 2
	// 3
}
