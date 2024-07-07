package itx_test

import (
	"fmt"
	"strings"

	"github.com/BooleanCat/go-functional/v2/it/filter"
	"github.com/BooleanCat/go-functional/v2/it/itx"
)

func ExampleLines() {
	for line := range itx.Lines(strings.NewReader("one\ntwo\nthree\n")) {
		fmt.Println(string(line))
	}
	// Output:
	// one
	// two
	// three
}

func ExampleLinesString() {
	reader := strings.NewReader("one\ntwo\n\nthree\n")

	fmt.Println(itx.LinesString(reader).Left().Exclude(filter.IsZero[string]).Collect())
	// Output: [one two three]
}
