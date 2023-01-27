package p2600

import (
	"fmt"
	"io"
)

func Solve2630(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	board := make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, n)
		for j := 0; j < n; j++ {
			_, _ = fmt.Fscan(reader, &board[i][j])
		}
	}

	w, b := cutBoard(board, 0, 0, n)
	_, _ = fmt.Fprintf(writer, "%d\n%d", w, b)
}

func cutBoard(board [][]int, x, y, size int) (int, int) {
	if size == 1 {
		if board[y][x] == 1 {
			return 0, 1
		} else {
			return 1, 0
		}
	}
	w0, b0 := cutBoard(board, x, y, size/2)
	w1, b1 := cutBoard(board, x+size/2, y, size/2)
	w2, b2 := cutBoard(board, x, y+size/2, size/2)
	w3, b3 := cutBoard(board, x+size/2, y+size/2, size/2)

	w := w0 + w1 + w2 + w3
	if w == 0 {
		return 0, 1
	}

	b := b0 + b1 + b2 + b3
	if b == 0 {
		return 1, 0
	}
	return w, b
}
