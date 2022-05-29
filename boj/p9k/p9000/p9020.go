package p9000

import (
	"fmt"
	"strconv"
)

func Solve9020(scanner Scanner, writer Writer) {
	sieveOfEratosthenes := buildSieveOfEratosthenes()

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())

		res0, res1 := findGoldbachPartitions(v, sieveOfEratosthenes)
		_, _ = fmt.Fprintln(writer, res0, res1)
	}
}

func findGoldbachPartitions(n int, sieveOfEratosthenes [maxLen]bool) (int, int) {
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

const maxLen = 10000 + 1

func buildSieveOfEratosthenes() [maxLen]bool {
	sieveOfEratosthenes := [maxLen]bool{}
	for i := 2; i*i < maxLen; i++ {
		if !sieveOfEratosthenes[i] {
			for j := i * i; j < maxLen; j += i {
				sieveOfEratosthenes[j] = true
			}
		}
	}
	sieveOfEratosthenes[0] = true
	sieveOfEratosthenes[1] = true
	return sieveOfEratosthenes
}
