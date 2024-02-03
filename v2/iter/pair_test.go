package iter_test

import (
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/iter"
)

func TestPairStringer(t *testing.T) {
	foo := map[string]interface{}{
		"text": "Random Text",
	}
	pair1 := iter.Pair[string, interface{}]{One: "1", Two: foo}
	pair2 := iter.Pair[int, interface{}]{One: 2, Two: pair1}
	pair3 := iter.Pair[interface{}, interface{}]{One: pair1, Two: pair2}

	assert.Equal(t, pair1.String(), "(1, map[text:Random Text])")
	assert.Equal(t, pair2.String(), "(2, (1, map[text:Random Text]))")
	assert.Equal(t, pair3.String(), "((1, map[text:Random Text]), (2, (1, map[text:Random Text])))")
}
