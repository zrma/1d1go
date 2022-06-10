package p2500

import (
	"fmt"
	"strconv"
)

const maxLen = 8 + 1

func Solve2580(scanner Scanner, writer Writer) {
	var blanks [][2]int
	board := [maxLen][maxLen]int{}
	for y := 0; y < maxLen; y++ {
		for x := 0; x < maxLen; x++ {
			scanner.Scan()
			board[y][x], _ = strconv.Atoi(scanner.Text())

			if board[y][x] == 0 {
				blanks = append(blanks, [2]int{y, x})
			}
		}
	}

	sudoku(board, blanks, 0, writer)
}

func sudoku(board [maxLen][maxLen]int, blanks [][2]int, depth int, writer Writer) bool {
	if depth == len(blanks) {
		printBoard(board, writer)
		return true
	}

	y, x := blanks[depth][0], blanks[depth][1]

	for i := 1; i <= maxLen; i++ {
		if isValid(board, y, x, i) {
			board[y][x] = i
			if sudoku(board, blanks, depth+1, writer) {
				return true
			}
			board[y][x] = 0
		}
	}
	return false
}

func isValid(board [maxLen][maxLen]int, y, x, v int) bool {
	for i := 0; i < maxLen; i++ {
		if board[y][i] == v || board[i][x] == v {
			return false
		}
	}

	for i := 0; i < maxLen; i++ {
		if board[y/3*3+i/3][x/3*3+i%3] == v {
			return false
		}
	}

	return true
}

func printBoard(board [maxLen][maxLen]int, writer Writer) {
	for y := 0; y < maxLen; y++ {
		for x := 0; x < maxLen; x++ {
			v := board[y][x]
			_, _ = fmt.Fprint(writer, v)
			if x != maxLen-1 {
				_, _ = fmt.Fprint(writer, " ")
			}
		}
		_, _ = fmt.Fprintln(writer)
	}
}
