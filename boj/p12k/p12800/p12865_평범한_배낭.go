package p12800

import (
	"fmt"

	"1d1go/boj/p12k/p12000"
	"1d1go/utils/integer"
)

func Solve12865(reader p12000.Reader, writer p12000.Writer) {
	var n, k int
	_, _ = fmt.Fscan(reader, &n, &k)

	arr := make([][2]int, n+1)
	for i := 1; i <= n; i++ {
		_, _ = fmt.Fscan(reader, &arr[i][0], &arr[i][1])
	}

	res := knapsack(arr, n, k)
	_, _ = fmt.Fprint(writer, res)
}

func knapsack(arr [][2]int, n, k int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			if j >= arr[i][0] {
				// arr[i][0] -> 이 물건을 넣으면 추가 될 무게
				// j - arr[i][0] -> 이 물건을 넣고 남는 무게
				// dp[i-1][j-arr[i][0]] -> [이 물건을 넣고 남는 무게] 로 채울 수 있는 가장 높은 가치
				// dp[i-1][j-arr[i][0]]+arr[i][1] -> 이 무게로 채울 수 있는 가장 큰 가치
				dp[i][j] = integer.Max(dp[i-1][j], dp[i-1][j-arr[i][0]]+arr[i][1])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][k]
}
