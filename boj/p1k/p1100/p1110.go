package p1100

func Solve1110(n int) int {
	count := 1
	for x := modify(n); x != n; x = modify(x) {
		count++
	}
	return count
}

func modify(n int) int {
	remain := n % 10
	sumOfDigit := (n / 10) + remain
	return remain*10 + sumOfDigit%10
}
