package it_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/it"
)

func ExampleCompact() {
	words := slices.Values([]string{"", "foo", "", "", "bar", ""})
	fmt.Println(slices.Collect(it.Compact(words)))
	// Output: [foo bar]
}

func TestCompactEmpty(t *testing.T) {
	t.Parallel()

	words := slices.Collect(it.Compact(it.Exhausted[string]()))
	assert.Empty[string](t, words)
}

func TestCompactOnlyEmpty(t *testing.T) {
	t.Parallel()

	words := slices.Values([]string{"", "", "", ""})
	assert.Empty[string](t, slices.Collect(it.Compact(words)))
}

func TestCompactOnlyNotEmpty(t *testing.T) {
	t.Parallel()

	words := slices.Values([]string{"foo", "bar"})
	assert.SliceEqual(t, slices.Collect(it.Compact(words)), []string{"foo", "bar"})
}

func TestCompactYieldFalse(t *testing.T) {
	t.Parallel()

	words := it.Compact(slices.Values([]string{"foo", "bar"}))

	words(func(string) bool {
		return false
	})
}
