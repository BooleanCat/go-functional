package option_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/option"
)

func ExampleOption_Unwrap() {
	fmt.Println(option.Some(4).Unwrap())
	// Output: 4
}

func ExampleOption_UnwrapOr() {
	fmt.Println(option.Some(4).UnwrapOr(3))
	fmt.Println(option.None[int]().UnwrapOr(3))
	// Output:
	// 4
	// 3
}

func ExampleOption_UnwrapOrElse() {
	fmt.Println(option.Some(4).UnwrapOrElse(func() int {
		return 3
	}))

	fmt.Println(option.None[int]().UnwrapOrElse(func() int {
		return 3
	}))

	// Output:
	// 4
	// 3
}

func ExampleOption_UnwrapOrZero() {
	fmt.Println(option.Some(4).UnwrapOrZero())
	fmt.Println(option.None[int]().UnwrapOrZero())

	// Output
	// 4
	// 0
}

func ExampleOption_IsSome() {
	fmt.Println(option.Some(4).IsSome())
	fmt.Println(option.None[int]().IsSome())

	// Output:
	// true
	// false
}

func ExampleOption_IsNone() {
	fmt.Println(option.Some(4).IsNone())
	fmt.Println(option.None[int]().IsNone())

	// Output:
	// false
	// true
}

func ExampleOption_Value() {
	value, ok := option.Some(4).Value()
	fmt.Println(value)
	fmt.Println(ok)

	// Output:
	// 4
	// true
}

func ExampleMap() {
	double := option.Map(func(x int) int {
		return x * 2
	})

	fmt.Println(double(option.None[int]()))
	fmt.Println(double(option.Some(25)))

	// Output:
	// None
	// Some(50)
}

func TestSomeStringer(t *testing.T) {
	assert.Equal(t, fmt.Sprintf("%s", option.Some("foo")), "Some(foo)") //nolint:gosimple
	assert.Equal(t, fmt.Sprintf("%s", option.Some(42)), "Some(42)")     //nolint:gosimple
}

func TestNoneStringer(t *testing.T) {
	assert.Equal(t, fmt.Sprintf("%s", option.None[string]()), "None") //nolint:gosimple
}

func TestSomeUnwrap(t *testing.T) {
	assert.Equal(t, option.Some(42).Unwrap(), 42)
}

func TestNoneUnwrap(t *testing.T) {
	defer func() {
		assert.Equal(t, fmt.Sprint(recover()), "called `Option.Unwrap()` on a `None` value")
	}()

	option.None[string]().Unwrap()
	t.Error("did not panic")
}

func TestSomeUnwrapOr(t *testing.T) {
	assert.Equal(t, option.Some(42).UnwrapOr(3), 42)
}

func TestNoneUnwrapOr(t *testing.T) {
	assert.Equal(t, option.None[int]().UnwrapOr(3), 3)
}

func TestSomeUnwrapOrElse(t *testing.T) {
	assert.Equal(t, option.Some(42).UnwrapOrElse(func() int { return 41 }), 42)
}

func TestNoneUnwrapOrElse(t *testing.T) {
	assert.Equal(t, option.None[int]().UnwrapOrElse(func() int { return 41 }), 41)
}

func TestSomeUnwrapOrZero(t *testing.T) {
	assert.Equal(t, option.Some(42).UnwrapOrZero(), 42)
}

func TestNoneUnwrapOrZero(t *testing.T) {
	assert.Equal(t, option.None[int]().UnwrapOrZero(), 0)
}

func TestIsSome(t *testing.T) {
	assert.True(t, option.Some(42).IsSome())
	assert.False(t, option.None[int]().IsSome())
}

func TestIsNone(t *testing.T) {
	assert.False(t, option.Some(42).IsNone())
	assert.True(t, option.None[int]().IsNone())
}

func TestSomeValue(t *testing.T) {
	value, ok := option.Some(42).Value()
	assert.Equal(t, value, 42)
	assert.True(t, ok)
}

func TestNoneValue(t *testing.T) {
	value, ok := option.None[int]().Value()
	assert.Equal(t, value, 0)
	assert.False(t, ok)
}

func TestMapSome(t *testing.T) {
	fn := func(x int) bool {
		return x > 50
	}

	optFn := option.Map(fn)

	resultWithSome := optFn(option.Some(40))

	unwrappedSome, ok := resultWithSome.Value()
	assert.True(t, ok)
	assert.False(t, unwrappedSome)
}

func TestMapNone(t *testing.T) {
	fn := func(x int) bool {
		t.Error("fn was called!")
		return x > 50
	}

	optFn := option.Map(fn)
	resultWithNone := optFn(option.None[int]())

	_, ok := resultWithNone.Value()
	assert.False(t, ok)
}
