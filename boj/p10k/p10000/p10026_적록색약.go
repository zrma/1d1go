package p10000

import (
	"fmt"
)

func Solve10026(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	board0 := make([][]byte, n)
	board1 := make([][]byte, n)
	for i := 0; i < n; i++ {
		board0[i] = make([]byte, n)
		board1[i] = make([]byte, n)

		var s string
		_, _ = fmt.Fscan(reader, &s)

		for j, c := range s {
			board0[i][j] = byte(c)
			board1[i][j] = byte(c)
			if board1[i][j] == 'G' {
				board1[i][j] = 'R'
			}
		}
	}

	var res0, res1 int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board0[i][j] != 0 {
				res0++
				dfs10026(board0, i, j, n)
			}
			if board1[i][j] != 0 {
				res1++
				dfs10026(board1, i, j, n)
			}
		}
	}

	_, _ = fmt.Fprint(writer, res0, res1)
}

func dfs10026(board [][]byte, i int, j int, n int) {
	curr := board[i][j]
	board[i][j] = 0

	if i-1 >= 0 && board[i-1][j] == curr {
		dfs10026(board, i-1, j, n)
	}

	if i+1 < n && board[i+1][j] == curr {
		dfs10026(board, i+1, j, n)
	}

	if j-1 >= 0 && board[i][j-1] == curr {
		dfs10026(board, i, j-1, n)
	}

	if j+1 < n && board[i][j+1] == curr {
		dfs10026(board, i, j+1, n)
	}
}
