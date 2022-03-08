package ops_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
	"github.com/BooleanCat/go-functional/iter/ops"
)

func ExampleAdd() {
	total := iter.Fold[int](iter.Lift([]int{1, 2, 3}), 0, ops.Add[int])

	fmt.Println(total)
	// Output: 6
}

func ExampleMultiply() {
	product := iter.Fold[int](iter.Lift([]int{3, 4, 5}), 2, ops.Multiply[int])

	fmt.Println(product)
	// Output: 120
}

func ExampleBitwiseAnd() {
	overlap := iter.Fold[int](iter.Lift([]int{5, 7, 13}), -1, ops.BitwiseAnd[int])

	fmt.Println(overlap)
	// Output: 5
}

func ExampleBitwiseOr() {
	union := iter.Fold[int](iter.Lift([]int{1, 2, 6}), 0, ops.BitwiseOr[int])

	fmt.Println(union)
	// Output: 7
}

func ExampleBitwiseXor() {
	result := iter.Fold[int](iter.Lift([]int{1, 2, 6}), 0, ops.BitwiseXor[int])

	fmt.Println(result)
	// Output: 5
}

func TestAdd(t *testing.T) {
	assert.Equal(t, ops.Add(5, 6), 11)
}

func TestMultiply(t *testing.T) {
	assert.Equal(t, ops.Multiply(3, 8), 24)
}

func TestBitwiseAnd(t *testing.T) {
	assert.Equal(t, ops.BitwiseAnd(6, 10), 2)
}

func TestBitwiseOr(t *testing.T) {
	assert.Equal(t, ops.BitwiseOr(6, 10), 14)
}

func TestBitwiseXor(t *testing.T) {
	assert.Equal(t, ops.BitwiseXor(5, 6), 3)
}
