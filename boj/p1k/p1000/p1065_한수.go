package p1000

import (
	"fmt"
	"io"
)

func Solve1065(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	var count int
	for i := 1; i <= n; i++ {
		if arithmeticProgression(i) {
			count++
		}
	}
	_, _ = fmt.Fprint(writer, count)
}

// arithmeticProgression args: n (1~1000)
func arithmeticProgression(n int) bool {
	if n < 100 {
		return true
	}
	if n > 999 {
		return false
	}
	i := n / 100
	j := n % 100 / 10
	k := n % 10
	return i-j == j-k
}
