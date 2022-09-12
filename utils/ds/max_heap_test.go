package ds

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMaxHeap(t *testing.T) {
	const size = 321
	h := NewMaxHeap(size)
	assert.NotNil(t, h)
	assert.Equal(t, 0, h.Size())

	assert.Equal(t, size, cap(h.data))

	{
		top, ok := h.Peek()
		assert.False(t, ok)
		assert.Equal(t, math.MinInt32, top)
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
		assert.Equal(t, i, top)
	}

	for i := 99; i > 1; i-- {
		v, ok := h.Pop()
		assert.True(t, ok)
		assert.Equal(t, i, v)
		assert.Equal(t, i-1, h.Size())

		top, ok := h.Peek()
		assert.True(t, ok)
		assert.Equal(t, i-1, top)
	}

	{
		v, ok := h.Pop()
		assert.True(t, ok)
		assert.Equal(t, 1, v)
		assert.Equal(t, 0, h.Size())

		top, ok := h.Peek()
		assert.False(t, ok)
		assert.Equal(t, math.MinInt32, top)

		v, ok = h.Pop()
		assert.False(t, ok)
		assert.Equal(t, 0, v)
	}

	for i := 99; i > 0; i-- {
		h.Push(i)
		assert.Equal(t, 100-i, h.Size())

		top, ok := h.Peek()
		assert.True(t, ok)
		assert.Equal(t, 99, top)
	}

	for i := 99; i > 1; i-- {
		v, ok := h.Pop()
		assert.True(t, ok)
		assert.Equal(t, i, v)
		assert.Equal(t, i-1, h.Size())

		top, ok := h.Peek()
		assert.True(t, ok)
		assert.Equal(t, i-1, top)
	}

	{
		v, ok := h.Pop()
		assert.True(t, ok)
		assert.Equal(t, 1, v)
		assert.Equal(t, 0, h.Size())

		top, ok := h.Peek()
		assert.False(t, ok)
		assert.Equal(t, math.MinInt32, top)

		v, ok = h.Pop()
		assert.False(t, ok)
		assert.Equal(t, 0, v)
	}
}
