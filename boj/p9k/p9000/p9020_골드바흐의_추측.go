package p9000

import (
	"fmt"

	"1d1go/utils/integer"
)

func Solve9020(reader Reader, writer Writer) {
	const maxLen = 10000 + 1
	sieveOfEratosthenes := integer.BuildSieveOfEratosthenes(maxLen)

	var n int
	_, _ = fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)

		res0, res1 := findGoldbachPartitions(v, sieveOfEratosthenes)
		_, _ = fmt.Fprintln(writer, res0, res1)
	}
}

func findGoldbachPartitions(n int, sieveOfEratosthenes []bool) (int, int) {
	for i := n / 2; i > 1; i-- {
		if sieveOfEratosthenes[i] {
			continue
		}
		j := n - i
		if !sieveOfEratosthenes[j] {
			return i, j
		}
	}
	return 0, n
}
