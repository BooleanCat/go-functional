package fifo_test

import (
	"testing"

	"github.com/BooleanCat/go-functional/v2/internal/assert"
	"github.com/BooleanCat/go-functional/v2/internal/fifo"
)

func TestFifo(t *testing.T) {
	queue := fifo.New[int, string]()

	queue.Enqueue(1, "one")
	queue.Enqueue(2, "two")
	queue.Enqueue(3, "three")

	for i := 1; i <= 3; i++ {
		left, ok := queue.DequeueLeft()
		assert.True(t, ok)
		assert.Equal(t, left, i)
	}

	for _, s := range []string{"one", "two", "three"} {
		right, ok := queue.DequeueRight()
		assert.True(t, ok)
		assert.Equal(t, right, s)
	}
}

func TestFifoEmpty(t *testing.T) {
	queue := fifo.New[int, string]()

	_, ok := queue.DequeueLeft()
	assert.False(t, ok)

	_, ok = queue.DequeueRight()
	assert.False(t, ok)
}

func TestFifoLeftRight(t *testing.T) {
	queue := fifo.New[int, string]()

	queue.Enqueue(1, "one")

	left, ok := queue.DequeueLeft()
	assert.True(t, ok)
	assert.Equal(t, left, 1)

	right, ok := queue.DequeueRight()
	assert.True(t, ok)
	assert.Equal(t, right, "one")
}
