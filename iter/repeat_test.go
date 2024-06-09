package iter_test

import (
	"fmt"
	it "iter"
	"testing"

	"github.com/BooleanCat/go-functional/v2/future/slices"
	"github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleRepeat() {
	numbers := slices.Collect(iter.Take(iter.Repeat(42), 2))

	fmt.Println(numbers)
	// Output: [42 42]
}

func TestRepeatTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := it.Pull(iter.Repeat(42))
	stop()
}
