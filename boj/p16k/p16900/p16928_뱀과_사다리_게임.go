package p16900

import (
	"fmt"
	"io"
)

func Solve16928(reader io.Reader, writer io.Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	const max = 100 + 1
	ladder := make([]int, max)
	board := make([]int, max)

	for i := 0; i < n; i++ {
		var a, b int
		_, _ = fmt.Fscan(reader, &a, &b)
		ladder[a] = b
	}
	for i := 0; i < m; i++ {
		var a, b int
		_, _ = fmt.Fscan(reader, &a, &b)
		ladder[a] = b
	}

	queue := make([]int, 0, max)
	queue = append(queue, 1)

	solve16928(ladder, queue, board, writer)
}

func solve16928(ladder, queue, board []int, writer io.Writer) {
	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]

		if q == 100 {
			break
		}

		for i := 6; i >= 1; i-- {
			next := q + i
			if next > 100 {
				continue
			}
			if ladder[next] > 0 {
				next = ladder[next]
			}
			if board[next] == 0 {
				board[next] = board[q] + 1
				queue = append(queue, next)
			}
		}
	}
	_, _ = fmt.Fprint(writer, board[100])
}
