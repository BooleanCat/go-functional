package iter_test

import (
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
)

func TestCollect(t *testing.T) {
	items := iter.Collect[int](iter.Take[int](iter.Count(), 5))
	assert.SliceEqual(t, items, []int{0, 1, 2, 3, 4})
}

func TestCollectEmpty(t *testing.T) {
	items := iter.Collect[int](iter.Take[int](iter.Count(), 0))
	assert.Empty(t, items)
}
