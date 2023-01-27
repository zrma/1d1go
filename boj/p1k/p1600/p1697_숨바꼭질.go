package p1600

import (
	"fmt"
	"io"
)

func Solve1697(reader io.Reader, writer io.Writer) {
	var n, k int
	_, _ = fmt.Fscan(reader, &n, &k)

	const max = 200000 + 1
	points := make([]int, max)

	queue := make([]int, 0, max)
	queue = append(queue, n)

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if p == k {
			break
		}

		cur := points[p]
		if p > 0 && points[p-1] == 0 {
			points[p-1] = cur + 1
			queue = append(queue, p-1)
		}
		if p < max-1 && points[p+1] == 0 {
			points[p+1] = cur + 1
			queue = append(queue, p+1)
		}
		if p*2 < max && points[p*2] == 0 {
			points[p*2] = cur + 1
			queue = append(queue, p*2)
		}
	}
	_, _ = fmt.Fprint(writer, points[k])
}
