package itx_test

import (
	"fmt"
	"maps"

	"github.com/BooleanCat/go-functional/v2/it/itx"
)

func ExampleOnce() {
	fmt.Println(itx.Once(42).Chain(itx.Once(43)).Collect())
	// Output: [42 43]
}

func ExampleOnce2() {
	numbers := maps.Collect(itx.Once2(1, 42).Chain(itx.Once2(2, 43)).Seq())
	fmt.Println(numbers[1])
	fmt.Println(numbers[2])

	// Output:
	// 42
	// 43
}
