package p17200

import (
	"fmt"
)

func Solve17298(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	res := make([]int, n)
	for i := range res {
		res[i] = -1
	}

	type entry struct {
		idx int
		val int
	}

	stack := make([]entry, 0, n)
	for i := 0; i < n; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)

		for len(stack) > 0 && stack[len(stack)-1].val < v {
			res[stack[len(stack)-1].idx] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, entry{idx: i, val: v})
	}

	for _, v := range res {
		_, _ = fmt.Fprintf(writer, "%d ", v)
	}
}
