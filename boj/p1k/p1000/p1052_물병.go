package p1000

import (
	"fmt"
)

func Solve1052(reader Reader, writer Writer) {
	var n, k int
	_, _ = fmt.Fscan(reader, &n, &k)

	if n <= k {
		_, _ = fmt.Fprint(writer, 0)
		return
	}

	added := 0
	for {
		bottle := countBottle(n + added)
		if bottle <= k {
			break
		}
		added++
	}
	_, _ = fmt.Fprint(writer, added)
}

func countBottle(n int) int {
	cnt := 0
	for n > 0 {
		if n%2 == 1 {
			cnt++
		}
		n /= 2
	}
	return cnt
}
