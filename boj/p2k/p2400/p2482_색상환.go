package p2400

import (
	"fmt"
	"io"
)

func Solve2482(reader io.Reader, writer io.Writer) {
	var n, k int
	_, _ = fmt.Fscan(reader, &n, &k)

	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, k+2)
		dp[i][0] = 1
		dp[i][1] = i
	}

	const mod = 1_000_000_003

	for i := 2; i <= n; i++ {
		for j := 2; j <= k; j++ {
			//  ... i-2  i-1  i  <- i  번째 색을 고르면
			//  ... i-2   x  [i]    i-1번째 색은 고를 수 없다. (i번째까지 포함해서 j개)
			// [... i-2]         <- (i-2)개 중에 j-1개롤 고르는 경우의 수 (i번째까지 포함해서 j개 참고)
			//
			//  ... i-2  i-1  i  <- i-1번째 색을 고르면
			// [... i-2  i-1] x     i  번째 색은 고를 수 없다. (i-1번째까지 포함해서 j개)
			// [... i-2  i-1]    <- (n-1)개 중에 j  개를 고르는 경우의 수 (i-1번째까지 포함해서 j개 참고)
			dp[i][j] = (dp[i-2][j-1] + dp[i-1][j]) % mod
		}
	}

	// 첫번째 색을 고를 경우, 두번째 색과 마지막 색은 고를 수 없다. (idx n-1, 0, 1)
	// (n-3) 개 중에 (k-1)개를 고르는 경우의 문제 (k개를 골라야 하는데 첫 번재 색이 1개 차지했으므로 k-1)
	//
	// 첫 번째 색을 안 고를 경우, n-1개 중에 k개를 고르는 경우의 문제
	res := (dp[n-3][k-1] + dp[n-1][k]) % mod
	_, _ = fmt.Fprint(writer, res)
}
