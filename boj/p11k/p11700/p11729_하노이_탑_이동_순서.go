package p11700

import (
	"fmt"
	"math"
	"strconv"
)

func Solve11729(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	step := int64(math.Pow(2, float64(n))) - 1
	_, _ = fmt.Fprintln(writer, step)

	from := 1
	to := 3
	other := 2
	hanoi(writer, n, from, to, other)

}

func hanoi(writer Writer, n int, from, to, other int) {
	if n == 1 {
		_, _ = fmt.Fprintln(writer, from, to)
		return
	}
	hanoi(writer, n-1, from, other, to)
	_, _ = fmt.Fprintln(writer, from, to)
	hanoi(writer, n-1, other, to, from)
}
