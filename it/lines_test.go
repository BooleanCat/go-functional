package it_test

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/it"
)

func ExampleLines() {
	buffer := strings.NewReader("one\ntwo\nthree\n")
	for line, err := range it.Lines(buffer) {
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println(string(line))
	}
	// Output:
	// one
	// two
	// three
}

func TestLinesError(t *testing.T) {
	t.Parallel()

	// Make string 66k bytes long.
	var longLine strings.Builder
	for i := 0; i < 66*1024; i++ {
		longLine.WriteByte('a')
	}

	buffer := strings.NewReader(longLine.String())
	for _, err := range it.Lines(buffer) {
		assert.True(t, err != nil)
	}
}

func TestLinesYieldsFalse(t *testing.T) {
	t.Parallel()

	buffer := strings.NewReader("one\ntwo\nthree\n")
	seq := it.Lines(buffer)

	seq(func(l []byte, e error) bool {
		return false
	})
}

func TestLinesYieldsFalseWithError(t *testing.T) {
	t.Parallel()

	// Make string 66k bytes long.
	var longLine strings.Builder
	for i := 0; i < 66*1024; i++ {
		longLine.WriteByte('a')
	}

	buffer := strings.NewReader(longLine.String())
	seq := it.Lines(buffer)

	seq(func(l []byte, e error) bool {
		return false
	})
}

type failSecondTime struct {
	count int
}

func (f *failSecondTime) Read(p []byte) (n int, err error) {
	if f.count == 0 {
		f.count++
		copy(p, []byte("o"))
		return 1, nil
	}

	if f.count == 1 {
		f.count++
		return 0, errors.New("read error")
	}

	return 0, io.EOF
}

var _ io.Reader = new(failSecondTime)

func TestLinesFailsLater(t *testing.T) {
	t.Parallel()

	var (
		count   int
		lastErr error
	)

	for _, err := range it.LinesString(new(failSecondTime)) {
		count++
		lastErr = err
	}

	assert.Equal(t, lastErr.Error(), "read error")
	assert.Equal(t, count, 2)
}

func ExampleLinesString() {
	buffer := strings.NewReader("one\ntwo\nthree\n")

	for line, err := range it.LinesString(buffer) {
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println(line)
	}
	// Output:
	// one
	// two
	// three
}
