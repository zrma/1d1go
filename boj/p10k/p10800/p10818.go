package p10800

func Solve10818(arr []int) [2]int {
	min := arr[0]
	max := arr[0]

	for _, n := range arr[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return [2]int{min, max}
}
