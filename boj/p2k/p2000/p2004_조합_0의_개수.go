package p2000

import (
	"fmt"
	"strconv"
)

func Solve2004(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	// nCm = n! / (m! * (n-m)!)
	countOfTwo := countNumInFactorial(n, 2) -
		countNumInFactorial(m, 2) -
		countNumInFactorial(n-m, 2)
	countOfFive := countNumInFactorial(n, 5) -
		countNumInFactorial(m, 5) -
		countNumInFactorial(n-m, 5)
	countOfTen := countOfTwo
	if countOfTen > countOfFive {
		countOfTen = countOfFive
	}
	_, _ = fmt.Fprint(writer, countOfTen)
}

func countNumInFactorial(n, target int) int {
	count := 0
	for i := target; i <= n; i *= target {
		count += n / i
	}
	return count
}
