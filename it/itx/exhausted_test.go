package itx_test

import (
	"fmt"

	"github.com/BooleanCat/go-functional/v2/it/itx"
)

func ExampleExhausted() {
	fmt.Println(len(itx.Exhausted[int]().Collect()))
	// Output: 0
}

func ExampleExhausted2() {
	fmt.Println(len(itx.Exhausted2[int, string]().Left().Collect()))
	// Output: 0
}
