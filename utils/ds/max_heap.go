package ds

import (
	"container/heap"
	"math"
)

func NewMaxHeap(size int) *MaxHeap {
	return &MaxHeap{pq: NewPriorityQueue(size)}
}

type MaxHeap struct {
	pq *PriorityQueue
}

func (h *MaxHeap) Push(v int) {
	heap.Push(h.pq, maxHeapItem{v})
}

func (h *MaxHeap) Peek() (int, bool) {
	if h.pq.Len() == 0 {
		return math.MinInt32, false
	}
	return h.pq.items[0].(maxHeapItem).v, true
}

func (h *MaxHeap) Pop() (int, bool) {
	if h.pq.Len() == 0 {
		return 0, false
	}

	return heap.Pop(h.pq).(maxHeapItem).v, true
}

func (h *MaxHeap) Size() int {
	return h.pq.Len()
}

var _ Priority = (*maxHeapItem)(nil)

type maxHeapItem struct{ v int }

func (m maxHeapItem) Priority() int { return -m.v }
