package lv0easy

import "math"

func rev(n int) int {
	var res int
	for n > 0 {
		res *= 10
		res += n % 10
		n /= 10
	}

	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}
	return res
}

func reverse(n int) int {
	negative := n < 0
	if negative {
		n = -n
		return -rev(n)
	}
	return rev(n)
}
