package result_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/result"
)

func ExampleResult_Unwrap() {
	fmt.Println(result.Ok(4).Unwrap())
	// Output: 4
}

func ExampleResult_UnwrapOr() {
	fmt.Println(result.Ok(4).UnwrapOr(3))
	fmt.Println(result.Err[int](errors.New("oops")).UnwrapOr(3))
	// Output:
	// 4
	// 3
}

func ExampleResult_UnwrapOrElse() {
	fmt.Println(result.Ok(4).UnwrapOrElse(func() int {
		return 3
	}))

	fmt.Println(result.Err[int](errors.New("oops")).UnwrapOrElse(func() int {
		return 3
	}))

	// Output:
	// 4
	// 3
}

func ExampleResult_UnwrapOrZero() {
	fmt.Println(result.Ok(4).UnwrapOrZero())
	fmt.Println(result.Err[int](errors.New("oops")).UnwrapOrZero())

	// Output
	// 4
	// 0
}

func ExampleResult_IsOk() {
	fmt.Println(result.Ok(4).IsOk())
	fmt.Println(result.Err[int](errors.New("oops")).IsOk())

	// Output:
	// true
	// false
}

func ExampleResult_IsErr() {
	fmt.Println(result.Ok(4).IsErr())
	fmt.Println(result.Err[int](errors.New("oops")).IsErr())

	// Output:
	// false
	// true
}

func ExampleResult_Value() {
	value, ok := result.Ok(4).Value()
	fmt.Println(value)
	fmt.Println(ok)

	// Output:
	// 4
	// <nil>
}

func ExampleResult_UnwrapErr() {
	err := result.Err[int](errors.New("oops")).UnwrapErr()
	fmt.Println(err)
	// Output: oops
}

func TestOkStringer(t *testing.T) {
	assert.Equal(t, fmt.Sprintf("%s", result.Ok(42)), "Ok(42)") //nolint:gosimple
}

func TestErrStringer(t *testing.T) {
	assert.Equal(t, fmt.Sprintf("%s", result.Err[int](errors.New("oops"))), "Err(oops)") //nolint:gosimple
}

func TestOkUnwrap(t *testing.T) {
	assert.Equal(t, result.Ok(42).Unwrap(), 42)
}

func TestErrUnwrap(t *testing.T) {
	defer func() {
		assert.Equal(t, fmt.Sprintf("%v", recover()), "called `Result.Unwrap()` on an `Err` value")
	}()

	result.Err[int](errors.New("oops")).Unwrap()
	t.Error("did not panic")
}

func TestOkUnwrapOr(t *testing.T) {
	assert.Equal(t, result.Ok(42).UnwrapOr(41), 42)
}

func TestErrUnwrapOr(t *testing.T) {
	assert.Equal(t, result.Err[int](errors.New("oops")).UnwrapOr(41), 41)
}

func TestOkUnwrapOrElse(t *testing.T) {
	assert.Equal(t, result.Ok(42).UnwrapOrElse(func() int { return 31 }), 42)
}

func TestErrUnwrapOrElse(t *testing.T) {
	assert.Equal(t, result.Err[int](errors.New("oops")).UnwrapOrElse(func() int { return 31 }), 31)
}

func TestOkUnwrapOrZero(t *testing.T) {
	assert.Equal(t, result.Ok(42).UnwrapOrZero(), 42)
}

func TestErrUnwrapOrZero(t *testing.T) {
	assert.Equal(t, result.Err[int](errors.New("oops")).UnwrapOrZero(), 0)
}

func TestIsOk(t *testing.T) {
	assert.True(t, result.Ok(42).IsOk())
	assert.False(t, result.Err[int](errors.New("oops")).IsOk())
}

func TestIsErr(t *testing.T) {
	assert.False(t, result.Ok(42).IsErr())
	assert.True(t, result.Err[int](errors.New("oops")).IsErr())
}

func TestOkValue(t *testing.T) {
	value, err := result.Ok(42).Value()
	assert.Equal(t, value, 42)
	assert.Nil(t, err)
}

func TestErrValue(t *testing.T) {
	value, err := result.Err[int](errors.New("oops")).Value()
	assert.Equal(t, value, 0)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "oops")
}

func TestErrUnwrapErr(t *testing.T) {
	err := result.Err[int](errors.New("oops")).UnwrapErr()
	assert.Equal(t, err.Error(), "oops")
}

func TestOkUnwrapErr(t *testing.T) {
	defer func() {
		assert.Equal(t, fmt.Sprintf("%v", recover()), "called `Result.UnwrapErr()` on an `Ok` value")
	}()

	_ = result.Ok(42).UnwrapErr()
	t.Error("did not panic")
}
