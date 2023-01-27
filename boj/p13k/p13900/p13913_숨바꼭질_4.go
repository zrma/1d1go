package p13900

import (
	"fmt"
	"io"
)

func Solve13913(reader io.Reader, writer io.Writer) {
	var n, k int
	_, _ = fmt.Fscan(reader, &n, &k)

	const max = 200000 + 1
	points := make([]int, max)
	prev := make([]int, max)

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
			prev[p-1] = p
			queue = append(queue, p-1)
		}
		if p < max-1 && points[p+1] == 0 {
			points[p+1] = cur + 1
			prev[p+1] = p
			queue = append(queue, p+1)
		}
		if p*2 < max && points[p*2] == 0 {
			points[p*2] = cur + 1
			prev[p*2] = p
			queue = append(queue, p*2)
		}
	}
	_, _ = fmt.Fprintln(writer, points[k])

	path := make([]int, 0, points[k]+1)
	for p := k; p != n; p = prev[p] {
		path = append(path, p)
	}
	path = append(path, n)

	for i := len(path) - 1; i >= 0; i-- {
		_, _ = fmt.Fprintf(writer, "%d ", path[i])
	}
}
