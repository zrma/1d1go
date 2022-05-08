package p10900

func Solve10950(arr2D [][2]int) []int {
	res := make([]int, len(arr2D))
	for i, arr := range arr2D {
		res[i] = arr[0] + arr[1]
	}
	return res
}
