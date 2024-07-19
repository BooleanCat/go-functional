package itx_test

import (
	"fmt"
	"strings"

	"github.com/BooleanCat/go-functional/v2/it/itx"
)

func ExampleIterator2_Left() {
	text := strings.NewReader("one\ntwo\nthree\n")

	fmt.Println(itx.LinesString(text).Left().Collect())
	// Output: [one two three]
}

func ExampleIterator2_Right() {
	for value := range itx.FromMap(map[int]string{1: "one"}).Right() {
		fmt.Println(value)
	}
	// Output: one
}
