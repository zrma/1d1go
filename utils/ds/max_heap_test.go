package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMaxHeap(t *testing.T) {
	const size = 321
	h := NewMaxHeap(size)
	assert.NotNil(t, h)
	assert.Equal(t, 0, h.Size())

	assert.Equal(t, size, cap(h.data))
	assert.Equal(t, minVal, h.Peek())
	assert.Equal(t, 0, h.Pop())

	for i := 1; i < 100; i++ {
		h.Push(i)
		assert.Equal(t, i, h.Size())
		assert.Equal(t, i, h.Peek())
	}

	for i := 99; i > 1; i-- {
		assert.Equal(t, i, h.Pop())
		assert.Equal(t, i-1, h.Size())
		assert.Equal(t, i-1, h.Peek())
	}

	assert.Equal(t, 1, h.Pop())
	assert.Equal(t, 0, h.Size())
	assert.Equal(t, minVal, h.Peek())
	assert.Equal(t, 0, h.Pop())

	for i := 99; i > 0; i-- {
		h.Push(i)
		assert.Equal(t, 100-i, h.Size())
		assert.Equal(t, 99, h.Peek())
	}

	for i := 99; i > 1; i-- {
		assert.Equal(t, i, h.Pop())
		assert.Equal(t, i-1, h.Size())
		assert.Equal(t, i-1, h.Peek())
	}

	assert.Equal(t, 1, h.Pop())
	assert.Equal(t, 0, h.Size())
	assert.Equal(t, minVal, h.Peek())
	assert.Equal(t, 0, h.Pop())
}
