package ds

import (
	"math"
)

func NewMinHeap(size int) *MinHeap {
	return &MinHeap{data: make([]int, 0, size)}
}

type MinHeap struct {
	data []int
}

func (h *MinHeap) Push(v int) {
	h.data = append(h.data, v)
	h.up(len(h.data) - 1)
}

func (h *MinHeap) Peek() (int, bool) {
	if len(h.data) == 0 {
		return math.MaxInt32, false
	}
	return h.data[0], true
}

func (h *MinHeap) Pop() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}

	v := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.down(0)

	return v, true
}

func (h *MinHeap) Size() int {
	return len(h.data)
}

func (h *MinHeap) up(cur int) {
	for cur > 0 {
		parent := (cur - 1) / 2
		if h.data[parent] < h.data[cur] {
			break
		}
		h.data[parent], h.data[cur] = h.data[cur], h.data[parent]
		cur = parent
	}
}

func (h *MinHeap) down(cur int) {
	for cur*2+1 < len(h.data) {
		child := cur*2 + 1
		if child+1 < len(h.data) && h.data[child+1] < h.data[child] {
			child++
		}
		if h.data[cur] <= h.data[child] {
			break
		}
		h.data[cur], h.data[child] = h.data[child], h.data[cur]
		cur = child
	}
}
