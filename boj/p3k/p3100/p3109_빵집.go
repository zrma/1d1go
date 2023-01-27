package p3100

import (
	"fmt"
	"io"
)

type gridState int

const (
	empty gridState = iota
	wall
	visited
)

func Solve3109(reader io.Reader, writer io.Writer) {
	var r, c int
	_, _ = fmt.Fscan(reader, &r, &c)

	grid := make([][]gridState, r)
	for row := range grid {
		var s string
		_, _ = fmt.Fscan(reader, &s)
		grid[row] = make([]gridState, c)
		for col := range grid[row] {
			c := s[col]
			if c == 'x' {
				grid[row][col] = wall
			}
		}
	}

	_, _ = fmt.Fprint(writer, solve3109(grid))
}

func solve3109(grid [][]gridState) int {
	res := 0
	for row := range grid {
		if dfs3109(grid, row, 0) {
			res++
		}
	}
	return res
}

func dfs3109(grid [][]gridState, row int, col int) bool {
	if col == len(grid[row])-1 {
		return true
	}

	if row > 0 && grid[row-1][col+1] == empty {
		grid[row-1][col+1] = visited
		if dfs3109(grid, row-1, col+1) {
			return true
		}
	}

	if grid[row][col+1] == empty {
		grid[row][col+1] = visited
		if dfs3109(grid, row, col+1) {
			return true
		}
	}

	if row < len(grid)-1 && grid[row+1][col+1] == empty {
		grid[row+1][col+1] = visited
		if dfs3109(grid, row+1, col+1) {
			return true
		}
	}

	return false
}
