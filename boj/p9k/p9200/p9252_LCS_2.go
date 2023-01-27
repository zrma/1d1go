package p9200

import (
	"fmt"
	"io"
)

func Solve9252(reader io.Reader, writer io.Writer) {
	var s0, s1 string
	_, _ = fmt.Fscan(reader, &s0, &s1)

	n0, n1 := len(s0), len(s1)
	dp := make([][]int, n0+1)
	for i := range dp {
		dp[i] = make([]int, n1+1)
	}

	LongestCommonSubsequence(n0, n1, s0, s1, dp)

	_, _ = fmt.Fprintln(writer, dp[n0][n1])

	res := make([]rune, 0, dp[n0][n1])
	for i, j := n0, n1; i > 0 && j > 0; {
		if s0[i-1] == s1[j-1] {
			res = append(res, rune(s0[i-1]))
			i--
			j--
		} else {
			if dp[i-1][j] > dp[i][j-1] {
				i--
			} else {
				j--
			}
		}
	}
	for i := len(res) - 1; i >= 0; i-- {
		_, _ = fmt.Fprintf(writer, "%c", res[i])
	}
}
