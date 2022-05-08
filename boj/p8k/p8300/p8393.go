package p8300

func Solve8393(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func Solve8393AP(n int) int {
	return n * (n + 1) / 2
}
