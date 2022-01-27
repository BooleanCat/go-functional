package filters_test

import (
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
	"github.com/BooleanCat/go-functional/iter/filters"
)

func TestIsZero(t *testing.T) {
	items := iter.Exclude[int](iter.Lift([]int{1, 2, 3, 0, 4}), filters.IsZero[int])
	assert.SliceEqual(t, iter.Collect[int](items), []int{1, 2, 3, 4})
}

func TestGreaterThan(t *testing.T) {
	items := iter.Filter[int](iter.Lift([]int{1, 2, 3, 4, 5, 1}), filters.GreaterThan(2))
	assert.SliceEqual(t, iter.Collect[int](items), []int{3, 4, 5})
}

func TestLessThan(t *testing.T) {
	items := iter.Filter[int](iter.Lift([]int{1, 2, 3, 4, 5, 1}), filters.LessThan(2))
	assert.SliceEqual(t, iter.Collect[int](items), []int{1, 1})
}
