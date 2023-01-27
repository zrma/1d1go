package p1600

import (
	"fmt"
	"io"

	"1d1go/utils/integer"
)

func Solve1644(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	sieveOfEratosthenes := integer.BuildSieveOfEratosthenes(n + 1)

	primes := make([]int, 0, n)
	for i := 2; i <= n; i++ {
		if !sieveOfEratosthenes[i] {
			primes = append(primes, i)
		}
	}

	begin, end := 0, 0
	res := 0
	for end <= len(primes) {
		sum := 0
		for i := begin; i < end; i++ {
			sum += primes[i]
		}
		if sum == n {
			res++
			end++
		} else if sum < n {
			end++
		} else {
			begin++
		}
	}

	_, _ = fmt.Fprint(writer, res)
}
