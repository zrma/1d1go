package p1900

import (
	"fmt"
)

func Solve1929(reader Reader, writer Writer) {
	var min, max int
	_, _ = fmt.Fscan(reader, &min, &max)

	const maxLen = 1000000
	sieveOfEratosthenes := [maxLen + 1]bool{}

	for i := 2; i*i <= max; i++ {
		if !sieveOfEratosthenes[i] {
			for j := i * i; j <= max; j += i {
				sieveOfEratosthenes[j] = true
			}
		}
	}
	sieveOfEratosthenes[0] = true
	sieveOfEratosthenes[1] = true

	for i := min; i <= max; i++ {
		if !sieveOfEratosthenes[i] {
			_, _ = fmt.Fprintln(writer, i)
		}
	}
}
