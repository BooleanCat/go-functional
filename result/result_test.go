package result_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/result"
)

func TestOkStringer(t *testing.T) {
	assert.Equal(t, fmt.Sprintf("%s", result.Ok(42)), "Ok(42)")
}

func TestErrStringer(t *testing.T) {
	assert.Equal(t, fmt.Sprintf("%s", result.Err[int](errors.New("oops"))), "Err(oops)")
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
