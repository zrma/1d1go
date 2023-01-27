package p14800

import (
	"fmt"
	"io"
)

func Solve14888(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([]int, n)
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	var plus, minus, times, divide int
	_, _ = fmt.Fscan(reader, &plus, &minus, &times, &divide)

	h := insertOpHelper{
		arr: arr,
		min: 1_000_000_000,
		max: -1_000_000_000,
	}
	h.insertOperator(arr[0], 1, plus, minus, times, divide)

	_, _ = fmt.Fprintln(writer, h.max)
	_, _ = fmt.Fprintln(writer, h.min)
}

type insertOpHelper struct {
	arr []int
	min int
	max int
}

func (h *insertOpHelper) insertOperator(res, depth, plus, minus, times, divide int) {
	if depth == len(h.arr) {
		if res < h.min {
			h.min = res
		}
		if res > h.max {
			h.max = res
		}
		return
	}

	if plus > 0 {
		h.insertOperator(res+h.arr[depth], depth+1, plus-1, minus, times, divide)
	}
	if minus > 0 {
		h.insertOperator(res-h.arr[depth], depth+1, plus, minus-1, times, divide)
	}
	if times > 0 {
		h.insertOperator(res*h.arr[depth], depth+1, plus, minus, times-1, divide)
	}
	if divide > 0 {
		h.insertOperator(res/h.arr[depth], depth+1, plus, minus, times, divide-1)
	}
}
