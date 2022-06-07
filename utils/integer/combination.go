package integer

func CombinationDP(n, r int, cache [][]int, modifier func(int) int) int {
	if n < r {
		return 0
	}
	if n == r {
		return 1
	}
	if r == 0 {
		return 1
	}
	if r == 1 {
		return n
	}

	if v := cache[n][r]; v != -1 {
		return v
	}

	res := CombinationDP(n-1, r-1, cache, modifier) + CombinationDP(n-1, r, cache, modifier)
	res = modifier(res)
	cache[n][r] = res

	return res
}

func BuildCache2DArr(maxLen, initVal int) [][]int {
	cache := make([][]int, maxLen)
	for i := range cache {
		cache[i] = make([]int, maxLen)
		for j := range cache[i] {
			cache[i][j] = initVal
		}
	}
	return cache
}
