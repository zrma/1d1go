package ds

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMinHeap(t *testing.T) {
	const size = 123
	h := NewMinHeap(size)
	assert.NotNil(t, h)
	assert.Equal(t, 0, h.Size())

	{
		top, ok := h.Peek()
		assert.False(t, ok)
		assert.Equal(t, math.MaxInt32, top)
		assert.Equal(t, 0, h.Size())

		v, ok := h.Pop()
		assert.False(t, ok)
		assert.Equal(t, 0, v)
	}

	for i := 1; i < 100; i++ {
		h.Push(i)
		assert.Equal(t, i, h.Size())

		top, ok := h.Peek()
		assert.True(t, ok)
		assert.Equal(t, 1, top)
	}

	for i := 1; i < 99; i++ {
		v, ok := h.Pop()
		assert.True(t, ok)
		assert.Equal(t, i, v)
		assert.Equal(t, 99-i, h.Size())

		top, ok := h.Peek()
		assert.True(t, ok)
		assert.Equal(t, i+1, top)
	}

	{
		v, ok := h.Pop()
		assert.True(t, ok)
		assert.Equal(t, 99, v)
		assert.Equal(t, 0, h.Size())

		top, ok := h.Peek()
		assert.False(t, ok)
		assert.Equal(t, math.MaxInt32, top)

		v, ok = h.Pop()
		assert.False(t, ok)
		assert.Equal(t, 0, v)
	}

	for i := 99; i > 0; i-- {
		h.Push(i)
		assert.Equal(t, 100-i, h.Size())

		top, ok := h.Peek()
		assert.True(t, ok)
		assert.Equal(t, i, top)
	}

	for i := 1; i < 99; i++ {
		v, ok := h.Pop()
		assert.True(t, ok)
		assert.Equal(t, i, v)
		assert.Equal(t, 99-i, h.Size())

		top, ok := h.Peek()
		assert.True(t, ok)
		assert.Equal(t, i+1, top)
	}

	{
		v, ok := h.Pop()
		assert.True(t, ok)
		assert.Equal(t, 99, v)
		assert.Equal(t, 0, h.Size())

		top, ok := h.Peek()
		assert.False(t, ok)
		assert.Equal(t, math.MaxInt32, top)

		v, ok = h.Pop()
		assert.False(t, ok)
		assert.Equal(t, 0, v)
	}
}
