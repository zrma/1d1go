package p8300

import (
	"fmt"
	"strconv"
)

func Solve8393(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}

	_, _ = fmt.Fprint(writer, sum)
}

func Solve8393AP(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	sum := n * (n + 1) / 2
	_, _ = fmt.Fprint(writer, sum)
}
