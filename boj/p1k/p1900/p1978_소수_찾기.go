package p1900

import (
	"fmt"
	"strconv"
)

func Solve1978(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	cnt := 0
	for i := 0; i < n; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		if isPrimeNumber(v) {
			cnt++
		}
	}
	_, _ = fmt.Fprint(writer, cnt)
}

func isPrimeNumber(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
