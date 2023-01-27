package p2200

import (
	"fmt"
	"io"
)

func Solve2292(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

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
