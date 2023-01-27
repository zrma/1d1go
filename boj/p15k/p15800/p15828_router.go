package p15800

import (
	"fmt"
	"io"
)

func Solve15828(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	queue := make([]int, 0, n)
	for {
		var v int
		_, _ = fmt.Fscan(reader, &v)

		if v == -1 {
			break
		}

		if v == 0 {
			queue = queue[1:]
			continue
		}

		if len(queue) == n {
			continue
		}

		queue = append(queue, v)
	}

	if len(queue) == 0 {
		_, _ = fmt.Fprint(writer, "empty")
		return
	}

	for _, v := range queue {
		_, _ = fmt.Fprint(writer, v, " ")
	}
}
