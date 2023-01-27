package p1000

import (
	"fmt"
	"io"
	"math/big"
)

func Solve1086(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([]string, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	var k int
	_, _ = fmt.Fscan(reader, &k)

	dp := make([][]int64, 1<<n)
	for i := range dp {
		dp[i] = make([]int64, k)
	}
	dp[0][0] = 1

	// x = a * 10^(len(b)) + b
	// x % k = (a * 10^(len(b)) + b) % k
	// x % k = (a * 10^(len(b)) % k + b % k
	// x % k = (a % k) * 10^(len(b)) + (b % k)
	// x % k = (a % k) * (10^(len(b)) % k) + (b % k)

	const maxLen = 50 + 1

	// 10^(len(b)) % k = 10^(len(b) - 1) * 10 % k
	cache0 := make([]int, maxLen)
	cache0[0] = 1 % k
	for i := 1; i < maxLen; i++ {
		cache0[i] = (cache0[i-1] * 10) % k
	}

	cache1 := make([]int, n)
	for i, s := range arr {
		cache1[i] = modOfStr(s, k)
	}

	for i := 0; i < 1<<n; i++ {
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				continue
			}
			for l := 0; l < k; l++ {
				dp[i|(1<<j)][(l*cache0[len(arr[j])]+cache1[j])%k] += dp[i][l]
			}
		}
	}

	denominator := int64(1)
	for i := int64(2); i <= int64(n); i++ {
		denominator *= i
	}

	numerator := dp[(1<<n)-1][0]

	rat := big.NewRat(numerator, denominator)
	_, _ = fmt.Fprint(writer, rat.String())
}

func modOfStr(s string, k int) int {
	res := 0
	for _, c := range s {
		res = (res*10 + int(c-'0')) % k
	}
	return res
}
