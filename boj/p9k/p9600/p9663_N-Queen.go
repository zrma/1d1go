package p9600

import (
	"fmt"
)

func Solve9663(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	board := make([]int, n)
	res := nQueen(board, 0, n)

	_, _ = fmt.Fprint(writer, res)
}

func nQueen(board []int, depth, n int) int {
	if depth == n {
		return 1
	}

	res := 0
	for i := 0; i < n; i++ {
		if isValid(depth, i, board) {
			board[depth] = i
			res += nQueen(board, depth+1, n)
		}
	}

	return res
}

func isValid(depth, i int, board []int) bool {
	for j := 0; j < depth; j++ {
		if board[j] == i || board[j] == i+depth-j || board[j] == i-depth+j {
			return false
		}
	}

	return true
}
