package p1700

import (
	"fmt"
)

func Solve1780(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	board := make([][]int, n)
	for i := range board {
		board[i] = make([]int, n)
		for j := range board[i] {
			_, _ = fmt.Fscan(reader, &board[i][j])
		}
	}

	minus, zero, plus := numberOfPapers(board, 0, 0, n)
	_, _ = fmt.Fprintln(writer, minus)
	_, _ = fmt.Fprintln(writer, zero)
	_, _ = fmt.Fprintln(writer, plus)
}

func numberOfPapers(board [][]int, x, y, size int) (int, int, int) {
	if size == 1 {
		switch board[y][x] {
		case -1:
			return 1, 0, 0
		case 0:
			return 0, 1, 0
		case 1:
			return 0, 0, 1
		}
	}
	minus0, zero0, plus0 := numberOfPapers(board, x, y, size/3)
	minus1, zero1, plus1 := numberOfPapers(board, x+size/3, y, size/3)
	minus2, zero2, plus2 := numberOfPapers(board, x+2*size/3, y, size/3)
	minus3, zero3, plus3 := numberOfPapers(board, x, y+size/3, size/3)
	minus4, zero4, plus4 := numberOfPapers(board, x+size/3, y+size/3, size/3)
	minus5, zero5, plus5 := numberOfPapers(board, x+2*size/3, y+size/3, size/3)
	minus6, zero6, plus6 := numberOfPapers(board, x, y+2*size/3, size/3)
	minus7, zero7, plus7 := numberOfPapers(board, x+size/3, y+2*size/3, size/3)
	minus8, zero8, plus8 := numberOfPapers(board, x+2*size/3, y+2*size/3, size/3)

	minus := minus0 + minus1 + minus2 + minus3 + minus4 + minus5 + minus6 + minus7 + minus8
	zero := zero0 + zero1 + zero2 + zero3 + zero4 + zero5 + zero6 + zero7 + zero8
	plus := plus0 + plus1 + plus2 + plus3 + plus4 + plus5 + plus6 + plus7 + plus8
	if zero == 0 && plus == 0 {
		return 1, 0, 0
	}
	if minus == 0 && plus == 0 {
		return 0, 1, 0
	}
	if minus == 0 && zero == 0 {
		return 0, 0, 1
	}

	return minus, zero, plus
}
