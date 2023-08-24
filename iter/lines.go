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

var _ Iterator[result.Result[[]byte]] = new(LinesIter)

// LinesString instantiates a [*LinesIter] with results converted to a string
// via a [*MapIter]. See [Lines] for more information.
func LinesString(r io.Reader) *MapIter[result.Result[[]byte], result.Result[string]] {
	iter := Lines(r)
	transform := func(line result.Result[[]byte]) result.Result[string] {
		if v, err := line.Value(); err != nil {
			return result.Err[string](err)
		} else {
			return result.Ok(string(v))
		}
	}

	return Map[result.Result[[]byte]](iter, transform)
}
