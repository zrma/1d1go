package p1900

import (
	"fmt"
)

func Solve1978(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	cnt := 0
	for i := 0; i < n; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)
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
