package itx_test

import (
	"fmt"

	"github.com/BooleanCat/go-functional/v2/it/itx"
)

func ExampleNaturalNumbers() {
	fmt.Println(itx.NaturalNumbers[int]().Take(4).Collect())
	// Output: [0 1 2 3]
}

func ExampleIntegers() {
	fmt.Println(itx.Integers[uint](0, 3, 1).Collect())
	// Output: [0 1 2]
}
