package iter_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
	"github.com/BooleanCat/go-functional/iter/ops"
	"github.com/BooleanCat/go-functional/result"
)

func ExampleLinesString() {
	lines := iter.LinesString(bytes.NewBufferString("hello\nthere"))
	unwrapped := iter.Map[result.Result[string]](lines, ops.UnwrapResult[string])

	fmt.Println(iter.Collect[string](unwrapped))
	// Output: [hello there]
}

func ExampleLines() {
	lines := iter.Lines(bytes.NewBufferString("hello\nthere"))
	unwrapped := iter.Map[result.Result[[]byte]](lines, ops.UnwrapResult[[]byte])

	fmt.Println(iter.Collect[[]byte](unwrapped))
	// Output: [[104 101 108 108 111] [116 104 101 114 101]]
}

func TestLines(t *testing.T) {
	file, err := os.Open("fixtures/lines.txt")
	assert.Nil(t, err)
	defer file.Close()

	lines := iter.Lines(file)

	assert.SliceEqual(t, lines.Next().Unwrap().Unwrap(), []byte("This is"))
	assert.SliceEqual(t, lines.Next().Unwrap().Unwrap(), []byte("a file"))
	assert.SliceEqual(t, lines.Next().Unwrap().Unwrap(), []byte("with"))
	assert.SliceEqual(t, lines.Next().Unwrap().Unwrap(), []byte("a trailing newline"))
	assert.SliceEqual(t, lines.Next().Unwrap().Unwrap(), []byte(""))
	assert.True(t, lines.Next().IsNone())
}

func TestLinesEmpty(t *testing.T) {
	lines := iter.Lines(new(bytes.Buffer))

	assert.SliceEqual(t, lines.Next().Unwrap().Unwrap(), []byte(""))
	assert.True(t, lines.Next().IsNone())
}

type readResult struct {
	content []byte
	err     error
}

type fakeReader struct {
	results []readResult
	index   int
}

func newFakeReader(results ...readResult) *fakeReader {
	return &fakeReader{results, 0}
}

func (r *fakeReader) Read(b []byte) (int, error) {
	r.index += 1
	if r.results[r.index-1].err != nil {
		return 0, r.results[r.index-1].err
	}
	n := copy(b, r.results[r.index-1].content)
	return n, nil
}

var _ io.Reader = new(fakeReader)

func TestLinesFailure(t *testing.T) {
	reader := newFakeReader(readResult{make([]byte, 0), errors.New("oops")})

	_, err := iter.Lines(reader).Next().Unwrap().Value()
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "read line: oops")
}

func TestLinesFailureLater(t *testing.T) {
	lines := iter.Lines(newFakeReader(
		readResult{[]byte("hello\n"), nil},
		readResult{make([]byte, 0), errors.New("oops")},
	))

	assert.SliceEqual(t, lines.Next().Unwrap().Unwrap(), []byte("hello"))
	assert.True(t, lines.Next().Unwrap().IsErr())
}

func TestLinesString(t *testing.T) {
	file, err := os.Open("fixtures/lines.txt")
	assert.Nil(t, err)
	defer file.Close()

	lines := iter.LinesString(file)

	assert.Equal(t, lines.Next().Unwrap().Unwrap(), "This is")
	assert.Equal(t, lines.Next().Unwrap().Unwrap(), "a file")
	assert.Equal(t, lines.Next().Unwrap().Unwrap(), "with")
	assert.Equal(t, lines.Next().Unwrap().Unwrap(), "a trailing newline")
	assert.Equal(t, lines.Next().Unwrap().Unwrap(), "")
	assert.True(t, lines.Next().IsNone())
}

func TestLinesStringEmpty(t *testing.T) {
	lines := iter.LinesString(new(bytes.Buffer))

	assert.Equal(t, lines.Next().Unwrap().Unwrap(), "")
	assert.True(t, lines.Next().IsNone())
}

func TestLinesStringFailure(t *testing.T) {
	lines := iter.LinesString(newFakeReader(readResult{make([]byte, 0), errors.New("oops")}))

	_, err := lines.Next().Unwrap().Value()
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "read line: oops")
}

func TestLinesStringFailureLater(t *testing.T) {
	lines := iter.LinesString(newFakeReader(
		readResult{[]byte("hello\n"), nil},
		readResult{make([]byte, 0), errors.New("oops")},
	))

	assert.Equal(t, lines.Next().Unwrap().Unwrap(), "hello")
	assert.True(t, lines.Next().Unwrap().IsErr())
}
