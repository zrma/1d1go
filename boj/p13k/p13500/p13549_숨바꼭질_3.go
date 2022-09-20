package p13300

import (
	"fmt"
)

func Solve13549(reader Reader, writer Writer) {
	var n, k int
	_, _ = fmt.Fscan(reader, &n, &k)

	const max = 200000 + 1
	points := make([]int, max)
	visits := make([]bool, max)

	queue := make([]int, 0, max)
	queue = append(queue, n)

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if p == k {
			break
		}

		cur := points[p]
		if p*2 < max && !visits[p*2] {
			points[p*2] = cur
			visits[p*2] = true
			queue = append([]int{p * 2}, queue...)
		}
		if p > 0 && !visits[p-1] {
			points[p-1] = cur + 1
			visits[p-1] = true
			queue = append(queue, p-1)
		}
		if p < max-1 && !visits[p+1] {
			points[p+1] = cur + 1
			visits[p+1] = true
			queue = append(queue, p+1)
		}
	}
	_, _ = fmt.Fprint(writer, points[k])
}
