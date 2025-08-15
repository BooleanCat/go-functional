package op_test

import (
	"errors"
	"fmt"
	"maps"
	"slices"
	"testing"

	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/BooleanCat/go-functional/v2/it/op"
)

func ExampleAdd() {
	fmt.Println(it.Fold(slices.Values([]int{1, 2, 3}), op.Add, 0))
	// Output: 6
}

func ExampleAdd_string() {
	fmt.Println(it.Fold(slices.Values([]string{"a", "b", "c"}), op.Add, ""))
	// Output: abc
}

func ExampleRef() {
	refs := slices.Collect(it.Map(slices.Values([]int{5, 6, 7}), op.Ref))
	fmt.Println(*refs[0], *refs[1], *refs[2])
	// Output: 5 6 7
}

func ExampleDeref() {
	intRef := func(a int) *int {
		return &a
	}

	values := slices.Values([]*int{intRef(4), intRef(5), intRef(6)})

	fmt.Println(slices.Collect(it.Map(values, op.Deref)))
	// Output: [4 5 6]
}

func ExampleMust() {
	numbers := maps.All(map[int]error{
		1: nil,
		2: nil,
		3: nil,
	})

	for number := range it.MapDown(numbers, op.Must) {
		fmt.Println(number)
	}
}

func TestMustPanic(t *testing.T) {
	t.Parallel()

	defer func() {
		r := recover()

		if r == nil {
			t.Error("expected panic")
		}
	}()

	slices.Collect(it.MapDown(maps.All(map[int]error{
		1: nil,
		2: nil,
		3: errors.New("error"),
	}), op.Must))
}
