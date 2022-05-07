package p2400

func Solve2480(arr [3]int) int {
	if arr[0] == arr[1] && arr[1] == arr[2] {
		return 10_000 + arr[0]*1_000
	}
	if arr[0] == arr[1] || arr[0] == arr[2] {
		return 1_000 + arr[0]*100
	}
	if arr[1] == arr[2] {
		return 1_000 + arr[1]*100
	}
	var max = arr[0]
	if arr[1] > max {
		max = arr[1]
	}
	if arr[2] > max {
		max = arr[2]
	}
	return max * 100
}
