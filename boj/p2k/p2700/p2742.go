package p2700

func Solve2742(n int) []int {
	res := make([]int, n)
	for i := n; i > 0; i-- {
		res[n-i] = i
	}
	return res
}
