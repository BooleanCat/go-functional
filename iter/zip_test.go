package iter_test

import (
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
)

func TestZip(t *testing.T) {
	isEven := func(a int) bool { return a%2 == 0 }
	evens := iter.Filter[int](iter.Count(), isEven)
	odds := iter.Exclude[int](iter.Count(), isEven)

	zipped := iter.Collect[iter.Tuple[int, int]](
		iter.Take[iter.Tuple[int, int]](iter.Zip[int, int](evens, odds), 3),
	)

	assert.SliceEqual(t, zipped, []iter.Tuple[int, int]{{0, 1}, {2, 3}, {4, 5}})
}
