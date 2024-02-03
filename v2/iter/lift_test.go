package iter_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/iter"
)

func ExampleLift() {
	for i := range iter.Lift([]int{1, 2}) {
		fmt.Println(i)
	}

	// Output:
	// 1
	// 2
}

func TestLift(t *testing.T) {
	number := 0

	for i := range iter.Lift([]int{1, 2}) {
		assert.Equal(t, number+1, i)
		number++
	}
}

func TestLiftEmpty(t *testing.T) {
	for _ = range iter.Lift([]int{}) {
		t.Error("expected no iteration")
	}
}

func ExampleLiftHashMap() {
	pairs := iter.LiftHashMap(map[string]int{"one": 1, "two": 2}).Collect()

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].One < pairs[j].One
	})

	fmt.Println(pairs)
	// Output: [(one, 1) (two, 2)]
}

func TestLiftHashMap(t *testing.T) {
	pairs := iter.LiftHashMap(map[string]int{"one": 1, "two": 2}).Collect()

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].One < pairs[j].One
	})

	assert.SliceEqual(t, []iter.Pair[string, int]{{"one", 1}, {"two", 2}}, pairs)
}

func TestLiftHashMapEmpty(t *testing.T) {
	for _ = range iter.LiftHashMap(map[string]int{}) {
		t.Error("expected no iteration")
	}
}
