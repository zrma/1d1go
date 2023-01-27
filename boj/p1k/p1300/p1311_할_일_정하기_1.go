package p1300

import (
	"fmt"
	"io"

	hungarian "github.com/oddg/hungarian-algorithm"
)

func Solve1311(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	var costs = make([][]int, n)
	for i := 0; i < n; i++ {
		costs[i] = make([]int, n)
		for j := 0; j < n; j++ {
			_, _ = fmt.Fscan(reader, &costs[i][j])
		}
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 1<<n)
	}

	res := solve1311WithDFS(0, 0, n, costs, dp)
	_, _ = fmt.Fprint(writer, res)
}

func solve1311WithDFS(cur, flag, n int, costs, dp [][]int) int {
	if cur == n {
		return 0
	}

	if dp[cur][flag] != 0 {
		return dp[cur][flag]
	}

	res := 1_234_567_891
	for i := 0; i < n; i++ {
		if flag&(1<<i) != 0 {
			continue
		}

		if cost := solve1311WithDFS(cur+1, flag|(1<<i), n, costs, dp) + costs[cur][i]; cost < res {
			res = cost
		}
	}

	dp[cur][flag] = res
	return res
}

// Solve1311WithHungarianAlgorithm solves the problem with Hungarian algorithm.
// TODO - 학습할 거리
// NOTE - https://en.wikipedia.org/wiki/Hungarian_algorithm
// NOTE - https://github.com/oddg/hungarian-algorithm
// NOTE - https://www.acmicpc.net/problem/14216
func Solve1311WithHungarianAlgorithm(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	var costs = make([][]int, n)
	for i := 0; i < n; i++ {
		costs[i] = make([]int, n)
		for j := 0; j < n; j++ {
			_, _ = fmt.Fscan(reader, &costs[i][j])
		}
	}

	match, _ := hungarian.Solve(costs)

	res := 0
	for row, col := range match {
		res += costs[row][col]
	}
	_, _ = fmt.Fprint(writer, res)
}
