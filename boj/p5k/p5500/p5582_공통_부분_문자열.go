package p5500

import (
	"fmt"
	"io"
)

func Solve5582(reader io.Reader, writer io.Writer) {
	var s0, s1 string
	_, _ = fmt.Fscan(reader, &s0, &s1)

	s0 = " " + s0
	s1 = " " + s1

	n0, n1 := len(s0), len(s1)
	dp := make([][]int, n0)
	for i := range dp {
		dp[i] = make([]int, n1)
	}

	LongestCommonSubstring(n0, n1, s0, s1, dp)

	max := 0
	for i := 1; i < n0; i++ {
		for j := 1; j < n1; j++ {
			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
	}

	_, _ = fmt.Fprint(writer, max)
}

func LongestCommonSubstring(n0, n1 int, s0, s1 string, dp [][]int) {
	for i := 1; i < n0; i++ {
		for j := 1; j < n1; j++ {
			if s0[i] == s1[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
		}
	}
}
