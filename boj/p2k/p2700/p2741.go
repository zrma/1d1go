package p2700

func Solve2741(n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = i + 1
	}
	return res
}
