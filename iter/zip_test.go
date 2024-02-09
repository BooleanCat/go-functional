package iter_test

import (
	"fmt"
	it "iter"
	"testing"

	"github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleZip() {
	for left, right := range iter.Zip(iter.Lift([]int{1, 2, 3}), iter.Lift([]string{"one", "two", "three"})) {
		fmt.Println(left, right)
	}

	// Output:
	// 1 one
	// 2 two
	// 3 three
}

func TestZipEmpty(t *testing.T) {
	t.Parallel()

	for _, _ = range iter.Zip(iter.Lift([]int{}), iter.Lift([]string{})) {
		t.Error("unexpected")
	}
}

func TestZipTerminateEarly(t *testing.T) {
	t.Parallel()

	_, stop := it.Pull2(it.Seq2[int, string](iter.Zip(iter.Lift([]int{1, 2}), iter.Lift([]string{"one", "two"}))))
	stop()
}
