package iter_test

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/internal/fakes"
	"github.com/BooleanCat/go-functional/iter"
	"github.com/BooleanCat/go-functional/option"
	"github.com/BooleanCat/go-functional/result"
)

func ExampleCollectResults() {
	words := iter.CollectResults[string](iter.LinesString(bytes.NewBufferString("hello\nfriend")))

	fmt.Println(words)
	// Output: Ok([hello friend])
}

func TestCollectResults(t *testing.T) {
	words := iter.CollectResults[string](iter.LinesString(bytes.NewBufferString("hello\nthere")))
	assert.SliceEqual[string](t, words.Unwrap(), []string{"hello", "there"})
}

func TestCollectResultsEmpty(t *testing.T) {
	words := iter.CollectResults[string](iter.Exhausted[result.Result[string]]())
	assert.Empty[string](t, words.Unwrap())
}

func TestCollectResultsErr(t *testing.T) {
	delegate := new(fakes.Iterator[result.Result[int]])
	delegate.NextReturns(option.Some(result.Err[int](errors.New("oops"))))

	numbers := iter.CollectResults[int](delegate)
	assert.Equal(t, numbers.UnwrapErr().Error(), "oops")
}

func TestCollectResultsErrStopsConsuming(t *testing.T) {
	delegate := new(fakes.Iterator[result.Result[int]])
	delegate.NextReturnsOnCall(0, option.Some(result.Ok(42)))
	delegate.NextReturnsOnCall(1, option.Some(result.Err[int](errors.New("oops"))))
	delegate.NextReturnsOnCall(2, option.Some(result.Ok(43)))

	numbers := iter.CollectResults[int](delegate)
	assert.Equal(t, numbers.UnwrapErr().Error(), "oops")
	assert.Equal(t, delegate.Next().Unwrap(), result.Ok(43))
}
