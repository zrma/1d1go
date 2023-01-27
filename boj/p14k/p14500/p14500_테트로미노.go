package p14500

import (
	"fmt"
	"io"
)

func Solve14500(reader io.Reader, writer io.Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	board := make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, m)
		for j := 0; j < m; j++ {
			_, _ = fmt.Fscan(reader, &board[i][j])
		}
	}

	max := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			area := calcArea(board, i, j, n, m)
			if area > max {
				max = area
			}
		}
	}
	_, _ = fmt.Fprint(writer, max)
}

func calcArea(board [][]int, i, j, n, m int) int {
	max := 0
	for _, block := range blocks {
		row := len(block)
		col := len(block[0])

		if i+row > n || j+col > m {
			continue
		}

		sum := 0
		for r := 0; r < row; r++ {
			for c := 0; c < col; c++ {
				if block[r][c] == 1 {
					sum += board[i+r][j+c]
				}
			}
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

var blocks = [19][][]int{
	{
		{1, 1, 1, 1},
	},
	{
		{1},
		{1},
		{1},
		{1},
	},
	{
		{1, 1},
		{1, 1},
	},
	{
		{1, 1, 1},
		{1, 0, 0},
	},
	{
		{1, 1},
		{0, 1},
		{0, 1},
	},
	{
		{0, 0, 1},
		{1, 1, 1},
	},
	{
		{1, 0},
		{1, 0},
		{1, 1},
	},
	{
		{1, 1, 1},
		{0, 0, 1},
	},
	{
		{0, 1},
		{0, 1},
		{1, 1},
	},
	{
		{1, 0, 0},
		{1, 1, 1},
	},
	{
		{1, 1},
		{1, 0},
		{1, 0},
	},
	{
		{1, 1, 1},
		{0, 1, 0},
	},
	{
		{0, 1},
		{1, 1},
		{0, 1},
	},
	{
		{0, 1, 0},
		{1, 1, 1},
	},
	{
		{1, 0},
		{1, 1},
		{1, 0},
	},
	{
		{1, 0},
		{1, 1},
		{0, 1},
	},
	{
		{0, 1, 1},
		{1, 1, 0},
	},
	{
		{0, 1},
		{1, 1},
		{1, 0},
	},
	{
		{1, 1, 0},
		{0, 1, 1},
	},
}
