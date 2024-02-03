package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleZip() {
	zipped := iter.Zip(iter.Count().Take(2), iter.Repeat("Hi")).Collect()
	fmt.Println(zipped)
	// Output: [(0, Hi) (1, Hi)]
}

func TestZip(t *testing.T) {
	zipped := iter.Zip(iter.Count().Take(2), iter.Repeat("Hi").Take(2)).Collect()
	assert.SliceEqual(t, []iter.Pair[int, string]{{0, "Hi"}, {1, "Hi"}}, zipped)
}

func TestZipEmpty(t *testing.T) {
	zipped := iter.Zip(iter.Exhausted[int](), iter.Exhausted[int]()).Collect()
	assert.Empty[iter.Pair[int, int]](t, zipped)
}

func TestZipEmptyLeft(t *testing.T) {
	zipped := iter.Zip(iter.Exhausted[int](), iter.Repeat(1)).Collect()
	assert.Empty[iter.Pair[int, int]](t, zipped)
}

func TestZipEmptyRight(t *testing.T) {
	zipped := iter.Zip(iter.Repeat(1), iter.Exhausted[int]()).Collect()
	assert.Empty[iter.Pair[int, int]](t, zipped)
}
