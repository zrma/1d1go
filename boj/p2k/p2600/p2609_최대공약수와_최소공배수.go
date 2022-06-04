package p2600

import (
	"fmt"
	"strconv"
)

func Solve2609(scanner Scanner, writer Writer) {
	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())

	_, _ = fmt.Fprintln(writer, gcd(a, b))
	_, _ = fmt.Fprintln(writer, lcm(a, b))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
