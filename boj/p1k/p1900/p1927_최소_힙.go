package p1900

import (
	"fmt"
)

func Solve1927(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	h := &minHeap{data: make([]int, 0, n)}
	for i := 0; i < n; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)

		if v == 0 {
			_, _ = fmt.Fprintln(writer, h.pop())
			continue
		}
		h.push(v)
	}
}

type minHeap struct {
	data []int
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
	if len(h.data) == 0 {
		return 0
	}

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
