package ds

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue(t *testing.T) {
	const size = 10
	pq := NewPriorityQueue(10)
	assert.Equal(t, size, cap(pq.items))
	assert.Equal(t, 0, pq.Len())

	heap.Push(pq, pqItem(2))
	heap.Push(pq, pqItem(3))
	heap.Push(pq, pqItem(1))

	assert.Equal(t, 3, pq.Len())

	assert.Equal(t, 1, int(heap.Pop(pq).(pqItem)))
	assert.Equal(t, 2, pq.Len())

	assert.Equal(t, 2, int(heap.Pop(pq).(pqItem)))
	assert.Equal(t, 1, pq.Len())

	assert.Equal(t, 3, int(heap.Pop(pq).(pqItem)))
	assert.Equal(t, 0, pq.Len())
}

var _ Priority = (*pqItem)(nil)

type pqItem int

func (i pqItem) Priority() int { return int(i) }
