package utils_test

import (
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/internal/utils.go"
)

func TestMin(t *testing.T) {
	assert.Equal(t, utils.Min(1, 2), 1)
	assert.Equal(t, utils.Min(2, 1), 1)
	assert.Equal(t, utils.Min(1, 1), 1)
}
