package p1500

import (
	"fmt"
)

func Solve1520(reader Reader, writer Writer) {
	var m, n int
	_, _ = fmt.Fscan(reader, &m, &n)

	arr = make([][]int64, m)
	dp = make([][]int64, m)
	for i := range dp {
		arr[i] = make([]int64, n)
		dp[i] = make([]int64, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			_, _ = fmt.Fscan(reader, &arr[i][j])
		}
	}
	res := findLowerDFS(n-1, m-1, n, m)
	_, _ = fmt.Fprint(writer, res)
}

var (
	arr [][]int64
	dp  [][]int64

	dx = []int{0, 1, 0, -1}
	dy = []int{1, 0, -1, 0}
)

func findLowerDFS(x, y, n, m int) int64 {
	if x == 0 && y == 0 {
		return 1
	}
	if dp[y][x] != -1 {
		return dp[y][x]
	}

	dp[y][x] = 0
	for i := 0; i < 4; i++ {
		newX := x + dx[i]
		newY := y + dy[i]

		if newX < 0 || newX >= n || newY < 0 || newY >= m {
			continue
		}

		if arr[y][x] < arr[newY][newX] {
			dp[y][x] += findLowerDFS(newX, newY, n, m)
		}
	}
	return dp[y][x]
}
