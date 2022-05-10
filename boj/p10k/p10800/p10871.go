package p10800

func Solve10871(x int, arr []int) []int {
	var res []int
	for _, n := range arr {
		if n < x {
			res = append(res, n)
		}
	}
	return res
}
