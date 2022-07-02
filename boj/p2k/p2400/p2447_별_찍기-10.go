package p2400

import (
	"fmt"
)

func Solve2447(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	board := make([][]rune, n)
	for i := 0; i < n; i++ {
		board[i] = make([]rune, n)
		for j := 0; j < n; j++ {
			board[i][j] = ' '
		}
	}

	printStarRecur(board, n, n/3, 0, 0)

	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			_, _ = fmt.Fprint(writer, string(board[y][x]))
		}
		_, _ = fmt.Fprint(writer, "\n")
	}
}

func printStarRecur(board [][]rune, n, k, x, y int) {
	if k == 1 {
		board[y][x], board[y][x+1], board[y][x+2] = '*', '*', '*'
		board[y+1][x], board[y+1][x+2] = '*', '*'
		board[y+2][x], board[y+2][x+1], board[y+2][x+2] = '*', '*', '*'
		return
	}
	printStarRecur(board, n, k/3, x, y)
	printStarRecur(board, n, k/3, x+k, y)
	printStarRecur(board, n, k/3, x+k+k, y)

	printStarRecur(board, n, k/3, x, y+k)
	printStarRecur(board, n, k/3, x+k+k, y+k)

	printStarRecur(board, n, k/3, x, y+k+k)
	printStarRecur(board, n, k/3, x+k, y+k+k)
	printStarRecur(board, n, k/3, x+k+k, y+k+k)
}
