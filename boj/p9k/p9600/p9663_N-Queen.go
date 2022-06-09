package p9600

import (
	"fmt"
	"strconv"
)

func Solve9663(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

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

func isValid(idx, i int, board []int) bool {
	for j := 0; j < idx; j++ {
		if board[j] == i || board[j] == i+idx-j || board[j] == i-idx+j {
			return false
		}
	}

	return true
}
