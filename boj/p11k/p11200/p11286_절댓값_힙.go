package p11200

import (
	"fmt"
)

func Solve11286(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	h := &absHeap{data: make([]int, 0, n)}
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

type absHeap struct {
	data []int
}

func (h *absHeap) push(v int) {
	h.data = append(h.data, v)
	h.up(len(h.data) - 1)
}

func (h *absHeap) up(cur int) {
	for cur > 0 {
		parent := (cur - 1) / 2
		if abs(h.data[parent]) < abs(h.data[cur]) {
			break
		}
		if abs(h.data[parent]) == abs(h.data[cur]) && h.data[parent] < h.data[cur] {
			break
		}
		h.data[parent], h.data[cur] = h.data[cur], h.data[parent]
		cur = parent
	}
}

func (h *absHeap) pop() int {
	if len(h.data) == 0 {
		return 0
	}

	v := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.down(0)

	return v
}

func (h *absHeap) down(cur int) {
	for cur*2+1 < len(h.data) {
		child := cur*2 + 1
		if child+1 < len(h.data) && (abs(h.data[child+1]) < abs(h.data[child]) || (abs(h.data[child+1]) == abs(h.data[child]) && h.data[child+1] < h.data[child])) {
			child++
		}
		if abs(h.data[cur]) < abs(h.data[child]) {
			break
		}
		if abs(h.data[cur]) == abs(h.data[child]) && h.data[cur] < h.data[child] {
			break
		}
		h.data[child], h.data[cur] = h.data[cur], h.data[child]
		cur = child
	}
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
