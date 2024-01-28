package iter

import (
	"bufio"
	"bytes"
	"fmt"
	"io"

	"github.com/BooleanCat/go-functional/option"
	"github.com/BooleanCat/go-functional/result"
)

// LinesIter iterator, see [Lines].
type LinesIter struct {
	BaseIter[result.Result[[]byte]]
	r        *bufio.Reader
	finished bool
}

// Lines instantiates a [*LinesIter] that will yield each line from the provided
// [io.Reader].
//
// Be aware that since Read operations can fail, the result time of each item
// is wrapped in a [result.Result].
func Lines(r io.Reader) *LinesIter {
	iter := &LinesIter{r: bufio.NewReader(r)}
	iter.BaseIter = BaseIter[result.Result[[]byte]]{iter}
	return iter
}

// Next implements the [Iterator] interface.
func (iter *LinesIter) Next() option.Option[result.Result[[]byte]] {
	if iter.finished {
		return option.None[result.Result[[]byte]]()
	}

	content, err := iter.r.ReadBytes('\n')

	if err == io.EOF {
		iter.finished = true
		return option.Some(result.Ok(content))
	}

	if err != nil {
		iter.finished = true
		return option.Some(result.Err[[]byte](fmt.Errorf(`read line: %w`, err)))
	}

	return option.Some(result.Ok(bytes.TrimRight(content, "\r\n")))
}

// String implements the [fmt.Stringer] interface
func (iter LinesIter) String() string {
	return "Iterator<Lines, type=Result<[]byte>>"
}

// CollectResults is a convenience method for [CollectResults], providing this
// iterator as an argument.
func (iter *LinesIter) CollectResults() result.Result[[][]byte] {
	return CollectResults[[]byte](iter)
}

var (
	_ fmt.Stringer                    = new(LinesIter)
	_ Iterator[result.Result[[]byte]] = new(LinesIter)
)

type LinesStringIter struct {
	BaseIter[result.Result[string]]
	iter *LinesIter
}

// Next implements the [Iterator] interface.
func (iter *LinesStringIter) Next() option.Option[result.Result[string]] {
	if value, ok := iter.iter.Next().Value(); !ok {
		return option.None[result.Result[string]]()
	} else {
		if b, err := value.Value(); err != nil {
			return option.Some(result.Err[string](err))
		} else {
			return option.Some(result.Ok(string(b)))
		}
	}
}

// String implements the [fmt.Stringer] interface
func (iter LinesStringIter) String() string {
	return "Iterator<LinesString, type=Result<string>>"
}

// CollectResults is a convenience method for [CollectResults], providing this
// iterator as an argument.
func (iter *LinesStringIter) CollectResults() result.Result[[]string] {
	return CollectResults[string](iter)
}

// LinesString instantiates a [*LinesStringIter] that behaves like a
// [*LinesIter] except that it yields strings. See [LinesIter].
func LinesString(r io.Reader) *LinesStringIter {
	iter := &LinesStringIter{iter: Lines(r)}
	iter.BaseIter = BaseIter[result.Result[string]]{iter}
	return iter
}

var (
	_ fmt.Stringer                    = new(LinesStringIter)
	_ Iterator[result.Result[string]] = new(LinesStringIter)
)
