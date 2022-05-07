package lv0easy

func removeDuplicates(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	cnt := 1
	prev := arr[0]
	for _, n := range arr[1:] {
		if n != prev {
			arr[cnt] = n
			cnt++
			prev = n
		}
	}
	return cnt
}
