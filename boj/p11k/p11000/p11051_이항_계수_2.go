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

	cache := make(map[int]map[int]int)
	res := binomialCoefficientWithCache(n, k, cache)
	_, _ = fmt.Fprint(writer, res)
}

func binomialCoefficientWithCache(n, k int, cache map[int]map[int]int) int {
	if n < k {
		return 0
	}
	if n == k {
		return 1
	}
	if k == 0 {
		return 1
	}

	c, ok := cache[n]
	if !ok {
		c = make(map[int]int)
		cache[n] = c
	}
	if v, ok := cache[n][k]; ok {
		return v
	}

	res := binomialCoefficientWithCache(n-1, k-1, cache) + binomialCoefficientWithCache(n-1, k, cache)
	res = res % 10_007
	c[k] = res

	return res
}
