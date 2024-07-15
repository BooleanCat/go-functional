package itx_test

import (
	"fmt"
	"maps"

	"github.com/BooleanCat/go-functional/v2/it/itx"
)

func ExampleRepeat() {
	fmt.Println(itx.Repeat(1).Take(3).Collect())
	// Output: [1 1 1]
}

func ExampleRepeat2() {
	fmt.Println(maps.Collect(itx.Repeat2(1, 2).Take(5).Seq()))
	// Output: map[1:2]
}
