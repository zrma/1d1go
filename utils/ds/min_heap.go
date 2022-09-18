package ds

import (
	"container/heap"
	"math"
)

func NewMinHeap(size int) *MinHeap {
	return &MinHeap{pq: NewPriorityQueue(size)}
}

type MinHeap struct {
	pq *PriorityQueue
}

func (h *MinHeap) Push(v int) {
	heap.Push(h.pq, minHeapItem{v})
}

func (h *MinHeap) Peek() (int, bool) {
	if h.pq.Len() == 0 {
		return math.MaxInt32, false
	}
	return h.pq.items[0].(minHeapItem).v, true
}

func (h *MinHeap) Pop() (int, bool) {
	if h.pq.Len() == 0 {
		return 0, false
	}
	return heap.Pop(h.pq).(minHeapItem).v, true
}

func (h *MinHeap) Size() int {
	return h.pq.Len()
}

var _ Priority = (*minHeapItem)(nil)

type minHeapItem struct{ v int }

func (m minHeapItem) Priority() int { return m.v }
