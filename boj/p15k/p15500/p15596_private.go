package p15500

func Solve15596(arr []int) int {
	var res int
	for _, n := range arr {
		res += n
	}
	return res
}
