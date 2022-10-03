package p1900

import (
	"fmt"

	"1d1go/utils/integer"
)

func Solve1929(reader Reader, writer Writer) {
	var min, max int
	_, _ = fmt.Fscan(reader, &min, &max)

	const maxLen = 1000000
	sieveOfEratosthenes := integer.BuildSieveOfEratosthenes(maxLen)

	for i := min; i <= max; i++ {
		if !sieveOfEratosthenes[i] {
			_, _ = fmt.Fprintln(writer, i)
		}
	}
}
