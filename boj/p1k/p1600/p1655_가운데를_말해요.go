package p1600

import (
	"fmt"
)

func Solve1655(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	maxH := &maxHeap{data: make([]int, 0, n)}
	minH := &minHeap{data: make([]int, 0, n)}

	for i := 0; i < n; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)

		if maxH.size() == minH.size() {
			maxH.push(v)
		} else {
			minH.push(v)
		}

		if maxH.peek() > minH.peek() {
			maxP, minP := maxH.pop(), minH.pop()
			maxH.push(minP)
			minH.push(maxP)
		}
		_, _ = fmt.Fprintln(writer, maxH.peek())
	}
}

const (
	maxVal = 999_999
)

type minHeap struct {
	data []int
}

func (h *minHeap) peek() int {
	if len(h.data) == 0 {
		return maxVal
	}
	return h.data[0]
}

func (h *minHeap) size() int {
	return len(h.data)
}

func (h *minHeap) push(v int) {
	h.data = append(h.data, v)
	h.up(len(h.data) - 1)
}

func (h *minHeap) up(cur int) {
	for cur > 0 {
		parent := (cur - 1) / 2
		if h.data[parent] < h.data[cur] {
			break
		}
		h.data[parent], h.data[cur] = h.data[cur], h.data[parent]
		cur = parent
	}
}

func (h *minHeap) pop() int {
	v := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.down(0)

	return v
}

func (h *minHeap) down(cur int) {
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

type maxHeap struct {
	data []int
}

func (h *maxHeap) peek() int {
	return h.data[0]
}

func (h *maxHeap) size() int {
	return len(h.data)
}

func (h *maxHeap) push(v int) {
	h.data = append(h.data, v)
	h.up(len(h.data) - 1)
}

func (h *maxHeap) up(cur int) {
	for cur > 0 {
		parent := (cur - 1) / 2
		if h.data[parent] >= h.data[cur] {
			break
		}
		h.data[parent], h.data[cur] = h.data[cur], h.data[parent]
		cur = parent
	}
}

func (h *maxHeap) pop() int {
	v := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.down(0)

	return v
}

func (h *maxHeap) down(cur int) {
	for cur*2+1 < len(h.data) {
		child := cur*2 + 1
		if child+1 < len(h.data) && h.data[child+1] > h.data[child] {
			child++
		}
		if h.data[cur] >= h.data[child] {
			break
		}
		h.data[cur], h.data[child] = h.data[child], h.data[cur]
		cur = child
	}
}
