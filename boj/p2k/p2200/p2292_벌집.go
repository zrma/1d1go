package p2200

import (
	"fmt"
	"strconv"
)

func Solve2292(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	got := calcStepInHive(n)
	_, _ = fmt.Fprint(writer, got)
}

func calcStepInHive(n int) int {
	sum := 1
	step := 1
	for ; sum < n; step++ {
		sum += 6 * step
	}
	return step
}
