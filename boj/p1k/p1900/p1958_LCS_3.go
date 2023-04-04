package p1900

import (
	"fmt"
	"io"
)

func Solve1958(reader io.Reader, writer io.Writer) {
	var s0, s1, s2 string
	_, _ = fmt.Fscan(reader, &s0, &s1, &s2)

	n0, n1, n2 := len(s0), len(s1), len(s2)
	dp := make([][][]int, n0+1)
	for i := range dp {
		dp[i] = make([][]int, n1+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n2+1)
		}
	}

	LongestCommonSubsequence(n0, n1, n2, s0, s1, s2, dp)

	_, _ = fmt.Fprint(writer, dp[n0][n1][n2])
}

func LongestCommonSubsequence(n0, n1, n2 int, s0, s1, s2 string, dp [][][]int) {
	for i := 1; i <= n0; i++ {
		for j := 1; j <= n1; j++ {
			for k := 1; k <= n2; k++ {
				if s0[i-1] == s1[j-1] && s1[j-1] == s2[k-1] {
					dp[i][j][k] = dp[i-1][j-1][k-1] + 1
				} else {
					if dp[i-1][j][k] > dp[i][j-1][k] {
						dp[i][j][k] = dp[i-1][j][k]
					} else {
						dp[i][j][k] = dp[i][j-1][k]
					}
					if dp[i][j][k] < dp[i][j][k-1] {
						dp[i][j][k] = dp[i][j][k-1]
					}
				}
			}
		}
	}
}
