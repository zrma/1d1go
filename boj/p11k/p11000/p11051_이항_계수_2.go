package p11000

import (
	"fmt"
	"strconv"
)

func Solve11051(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	const maxLen = 1000 + 1
	cache := make([][]int, maxLen)
	for i := range cache {
		cache[i] = make([]int, maxLen)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	res := binomialCoefficientWithCache(n, k, cache)
	_, _ = fmt.Fprint(writer, res)
}

func binomialCoefficientWithCache(n, k int, cache [][]int) int {
	if n < k {
		return 0
	}
	if n == k {
		return 1
	}
	if k == 0 {
		return 1
	}

	if v := cache[n][k]; v != -1 {
		return v
	}

	res := binomialCoefficientWithCache(n-1, k-1, cache) + binomialCoefficientWithCache(n-1, k, cache)
	res = res % 10_007
	cache[n][k] = res

	return res
}
