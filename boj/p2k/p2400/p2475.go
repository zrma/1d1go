package p2400

import (
	"fmt"
	"strconv"
)

func Solve2475(scanner Scanner, writer Writer) {
	const n = 5

	sum := 0
	for i := 0; i < n; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		sum += v * v
	}
	res := sum % 10
	_, _ = fmt.Fprint(writer, res)
}
