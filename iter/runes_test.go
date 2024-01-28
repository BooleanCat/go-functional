package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
)

func ExampleRunes() {
	runes := iter.Runes("Hello, 世界!").Collect()

	fmt.Println(string(runes[7:9]))
	// Output: 世界
}

func ExampleRunes_slice() {
	runes := iter.Runes([]rune("Hello, 世界!")).Collect()

	fmt.Println(string(runes[7:9]))
	// Output: 世界
}

func ExampleRunesIter_String() {
	fmt.Println(iter.Runes("Hello, 世界!"))
	// Output: Iterator<Runes>
}

func TestRunes(t *testing.T) {
	runes := iter.Runes("Hello, 世界!").Collect()
	assert.SliceEqual[rune](t, runes, []rune{'H', 'e', 'l', 'l', 'o', ',', ' ', '世', '界', '!'})
}

func TestRunesSlice(t *testing.T) {
	runes := iter.Runes([]rune("Hello, 世界!")).Collect()
	assert.SliceEqual[rune](t, runes, []rune{'H', 'e', 'l', 'l', 'o', ',', ' ', '世', '界', '!'})
}

func TestRunesEmpty(t *testing.T) {
	runes := iter.Runes("").Collect()
	assert.Empty[rune](t, runes)
}

func TestRunesEmptySlice(t *testing.T) {
	runes := iter.Runes([]rune{}).Collect()
	assert.Empty[rune](t, runes)
}

func TestRunesIter_String(t *testing.T) {
	runes := iter.Runes("Hello, 世界!")
	assert.Equal(t, fmt.Sprintf("%s", runes), "Iterator<Runes>")  //nolint:gosimple
	assert.Equal(t, fmt.Sprintf("%s", *runes), "Iterator<Runes>") //nolint:gosimple
}
