package p2600

import (
	"fmt"
	"io"
)

func Solve2618(reader io.Reader, writer io.Writer) {
	var n, w int
	_, _ = fmt.Fscan(reader, &n, &w)

	stepsA := make([][2]int, w+1)
	stepsB := make([][2]int, w+1)

	for i := 1; i <= w; i++ {
		_, _ = fmt.Fscan(reader, &stepsA[i][0], &stepsA[i][1])
	}

	copy(stepsB, stepsA)

	stepsA[0] = [2]int{1, 1}
	stepsB[0] = [2]int{n, n}

	dp := make([][]int, w+1)
	for i := range dp {
		dp[i] = make([]int, w+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	maxDist := getMaxDistance(0, 0, stepsA, stepsB, dp)
	_, _ = fmt.Fprintln(writer, maxDist)

	res := step(0, 0, stepsA, stepsB, dp)
	for _, v := range res {
		_, _ = fmt.Fprintln(writer, v)
	}
}

func getMaxDistance(a, b int, stepA, stepB [][2]int, dp [][]int) int {
	if a == len(stepA)-1 || b == len(stepB)-1 {
		return 0
	}

	if dp[a][b] != -1 {
		return dp[a][b]
	}

	nextStep := max(a, b) + 1

	distA := dist(a, nextStep, stepA) + getMaxDistance(nextStep, b, stepA, stepB, dp)
	distB := dist(b, nextStep, stepB) + getMaxDistance(a, nextStep, stepA, stepB, dp)

	dp[a][b] = min(distA, distB)
	return dp[a][b]
}

func step(a int, b int, stepA, stepB [][2]int, dp [][]int) []int {
	if a == len(stepA)-1 || b == len(stepB)-1 {
		return []int{}
	}

	nextStep := max(a, b) + 1

	distA := dist(a, nextStep, stepA) + getMaxDistance(nextStep, b, stepA, stepB, dp)
	distB := dist(b, nextStep, stepB) + getMaxDistance(a, nextStep, stepA, stepB, dp)

	if distA < distB {
		return append([]int{1}, step(nextStep, b, stepA, stepB, dp)...)
	} else {
		return append([]int{2}, step(a, nextStep, stepA, stepB, dp)...)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dist(from int, to int, steps [][2]int) int {
	return abs(steps[from][0]-steps[to][0]) + abs(steps[from][1]-steps[to][1])
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
