package p1400

import (
	"fmt"
	"strconv"
)

func Solve1436(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	res := 1
	num := 666
	for {
		if res == n {
			break
		}
		num++
		if contains666(num) {
			res++
		}
	}
	_, _ = fmt.Fprint(writer, num)
}

func contains666(n int) bool {
	for n >= 666 {
		if n%10 == 6 && n/10%10 == 6 && n/100%10 == 6 {
			return true
		}
		n /= 10
	}
	return false
}
