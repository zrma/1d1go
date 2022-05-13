package p1100

func Solve1110(fmt InOut) {
	var n int
	_, _ = fmt.Scan(&n)

	count := 1
	for x := modify(n); x != n; x = modify(x) {
		count++
	}

	_, _ = fmt.Println(count)
}

func modify(n int) int {
	remain := n % 10
	sumOfDigit := (n / 10) + remain
	return remain*10 + sumOfDigit%10
}
