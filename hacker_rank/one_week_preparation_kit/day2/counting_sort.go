package day2

func countingSort(arr []int32) []int32 {
	res := make([]int32, 100)
	for _, v := range arr {
		res[v]++
	}
	return res
}
