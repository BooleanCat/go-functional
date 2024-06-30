package itx_test

import (
	"fmt"

	"github.com/BooleanCat/go-functional/v2/it/itx"
)

func ExampleCount() {
	fmt.Println(itx.Count[int]().Take(4).Collect())
	// Output: [0 1 2 3]
}
