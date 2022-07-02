package p2500

import (
	"fmt"
	"math"
)

func Solve2581(reader Reader, writer Writer) {
	var m, n int
	_, _ = fmt.Fscan(reader, &m, &n)

	min := math.MaxInt
	sum := 0
	for i := m; i <= n; i++ {
		if isPrimeNumber(i) {
			sum += i
			if i < min {
				min = i
			}
		}
	}
	if sum == 0 {
		_, _ = fmt.Fprint(writer, -1)
	} else {
		_, _ = fmt.Fprintln(writer, sum)
		_, _ = fmt.Fprint(writer, min)
	}
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
