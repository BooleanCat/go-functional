package itx_test

import (
	"fmt"
	"maps"

	"github.com/BooleanCat/go-functional/v2/it/itx"
)

func ExampleExhausted() {
	fmt.Println(len(itx.Exhausted[int]().Collect()))
	// Output: 0
}

func ExampleExhausted2() {
	fmt.Println(len(maps.Collect(itx.Exhausted2[int, string]().Seq())))
	// Output: 0
}
