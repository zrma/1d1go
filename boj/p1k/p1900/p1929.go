package p1900

import (
	"fmt"
	"strconv"
)

func Solve1929(scanner Scanner, writer Writer) {
	scanner.Scan()
	min, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	max, _ := strconv.Atoi(scanner.Text())

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
