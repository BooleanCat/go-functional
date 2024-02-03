package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/iter"
	"github.com/BooleanCat/go-functional/v2/iter/filter"
)

func ExampleFilter() {
	fmt.Println(iter.Filter(iter.Count(), filter.IsEven).Take(3).Collect())
	// Output: [0 2 4]
}

func ExampleFilter_method() {
	fmt.Println(iter.Count().Filter(filter.IsEven).Take(3).Collect())
	// Output: [0 2 4]
}

func TestFilter(t *testing.T) {
	numbers := iter.Count().Filter(filter.IsEven).Take(3).Collect()
	assert.SliceEqual(t, []int{0, 2, 4}, numbers)
}

func TestFilterEmpty(t *testing.T) {
	numbers := iter.Exhausted[int]().Filter(filter.IsEven).Collect()
	assert.Empty[int](t, numbers)
}

func ExampleExclude() {
	fmt.Println(iter.Exclude(iter.Count(), filter.IsEven).Take(3).Collect())
	// Output: [1 3 5]
}

func ExampleExclude_method() {
	fmt.Println(iter.Count().Exclude(filter.IsEven).Take(3).Collect())
	// Output: [1 3 5]
}

func TestExclude(t *testing.T) {
	numbers := iter.Count().Exclude(filter.IsEven).Take(3).Collect()
	assert.SliceEqual(t, []int{1, 3, 5}, numbers)
}

func TestExcludeEmpty(t *testing.T) {
	numbers := iter.Exhausted[int]().Exclude(filter.IsEven).Collect()
	assert.Empty[int](t, numbers)
}
