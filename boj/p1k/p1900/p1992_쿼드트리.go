package p1900

import (
	"fmt"
)

func Solve1992(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	board := make([]string, n)
	for i := range board {
		_, _ = fmt.Fscan(reader, &board[i])
	}
	res := quadTree(board, 0, 0, n)
	_, _ = fmt.Fprint(writer, res)
}

func quadTree(board []string, x, y, size int) string {
	if size == 1 {
		if board[y][x] == '1' {
			return "1"
		} else {
			return "0"
		}
	}
	s0 := quadTree(board, x, y, size/2)
	s1 := quadTree(board, x+size/2, y, size/2)
	s2 := quadTree(board, x, y+size/2, size/2)
	s3 := quadTree(board, x+size/2, y+size/2, size/2)
	s := fmt.Sprintf("%s%s%s%s", s0, s1, s2, s3)

	if s == "0000" {
		return "0"
	}
	if s == "1111" {
		return "1"
	}
	return fmt.Sprintf("(%s)", s)
}
