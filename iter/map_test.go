package iter_test

import (
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
)

func TestMap(t *testing.T) {
	double := func(a int) int { return a * 2 }
	items := iter.Collect[int](iter.Take[int](
		iter.Map[int](iter.Count(), double),
		4,
	))
	assert.SliceEqual(t, items, []int{0, 2, 4, 6})
}

func TestMapEmpty(t *testing.T) {
	double := func(a int) int { return a * 2 }
	items := iter.Collect[int](iter.Map[int](iter.Exhausted[int](), double))
	assert.Empty(t, items)
}
