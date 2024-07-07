package itx_test

import (
	"fmt"
	"iter"
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
	lines, errs := itx.LinesString(strings.NewReader("one\ntwo\n\nthree\n\n")).Unzip()

	_, stop := iter.Pull(iter.Seq[error](errs))
	stop()

	fmt.Println(lines.Exclude(filter.IsZero[string]).Collect())
	// Output: [one two three]
}
