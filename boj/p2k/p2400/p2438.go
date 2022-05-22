package p2400

import (
	"fmt"
	"strconv"
)

func Solve2438(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			_, _ = fmt.Fprint(writer, "*")
		}
		_, _ = fmt.Fprintln(writer)
	}
}
