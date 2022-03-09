package ops_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
	"github.com/BooleanCat/go-functional/iter/ops"
	"github.com/BooleanCat/go-functional/option"
	"github.com/BooleanCat/go-functional/result"
)

func ExampleUnwrapOption() {
	options := iter.Lift([]option.Option[int]{
		option.Some(4),
		option.Some(6),
		option.Some(-1),
	})

	numbers := iter.Map[option.Option[int]](options, ops.UnwrapOption[int])

	fmt.Println(iter.Collect[int](numbers))
	// Output: [4 6 -1]
}

func ExampleUnwrapResult() {
	results := iter.Lift([]result.Result[int]{
		result.Ok(1),
		result.Ok(3),
		result.Ok(-6),
	})

	numbers := iter.Map[result.Result[int]](results, ops.UnwrapResult[int])

	fmt.Println(iter.Collect[int](numbers))
	// Output: [1 3 -6]
}

func ExamplePassthrough() {
	numbers := iter.Map[int](iter.Lift([]int{1, 2, 3}), ops.Passthrough[int])

	fmt.Println(iter.Collect[int](numbers))
	// Output: [1 2 3]
}

func TestUnwrapOption(t *testing.T) {
	options := iter.Lift([]option.Option[int]{
		option.Some(4),
		option.Some(6),
		option.Some(-1),
	})

	integers := iter.Map[option.Option[int]](options, ops.UnwrapOption[int])

	assert.SliceEqual(t, iter.Collect[int](integers), []int{4, 6, -1})
}

func TestUnwrapOptionPanic(t *testing.T) {
	defer func() {
		assert.Equal(t, fmt.Sprint(recover()), "called `Option.Unwrap()` on a `None` value")
	}()

	iter.Collect[int](
		iter.Map[option.Option[int]](
			iter.Lift(
				[]option.Option[int]{option.None[int]()},
			),
			ops.UnwrapOption[int],
		),
	)

	t.Error("did not panic")
}

func TestUnwrapResult(t *testing.T) {
	results := iter.Lift([]result.Result[int]{
		result.Ok(4),
		result.Ok(6),
		result.Ok(-1),
	})

	integers := iter.Map[result.Result[int]](results, ops.UnwrapResult[int])

	assert.SliceEqual(t, iter.Collect[int](integers), []int{4, 6, -1})
}

func TestUnwrapResultPanic(t *testing.T) {
	defer func() {
		assert.Equal(t, fmt.Sprint(recover()), "called `Result.Unwrap()` on an `Err` value")
	}()

	iter.Collect[int](
		iter.Map[result.Result[int]](
			iter.Lift(
				[]result.Result[int]{result.Err[int](errors.New("oops"))},
			),
			ops.UnwrapResult[int],
		),
	)

	t.Error("did not panic")
}

func TestPassthrough(t *testing.T) {
	iter := iter.Map[int](iter.Lift([]int{1, 2}), ops.Passthrough[int])
	assert.Equal(t, iter.Next().Unwrap(), 1)
	assert.Equal(t, iter.Next().Unwrap(), 2)
	assert.True(t, iter.Next().IsNone())
}
