package p3000

import (
	"fmt"
	"strconv"
)

func Solve3036(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	arr0, _ := strconv.Atoi(scanner.Text())

	for i := 1; i < n; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())

		div := gcd(arr0, v)
		_, _ = fmt.Fprintf(writer, "%d/%d\n", arr0/div, v/div)
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
