package p9200

import (
	"fmt"

	"1d1go/utils/integer"
)

func Solve9251(scanner Scanner, writer Writer) {
	scanner.Scan()
	s0 := scanner.Text()
	scanner.Scan()
	s1 := scanner.Text()

	res := LongestCommonSubSeqLen(s0, s1)
	_, _ = fmt.Fprint(writer, res)
}

func LongestCommonSubSeqLen(s0, s1 string) int {
	n0, n1 := len(s0), len(s1)
	dp := make([][]int, n0+1)
	for i := range dp {
		dp[i] = make([]int, n1+1)
	}

	for i := 1; i <= n0; i++ {
		for j := 1; j <= n1; j++ {
			if s0[i-1] == s1[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = integer.Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n0][n1]
}
