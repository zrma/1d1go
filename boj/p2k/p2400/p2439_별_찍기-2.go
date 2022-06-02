package p2400

import (
	"fmt"
	"strconv"
)

func Solve2439(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j < n-1-i {
				_, _ = fmt.Fprint(writer, " ")
			} else {
				_, _ = fmt.Fprint(writer, "*")
			}
		}
		_, _ = fmt.Fprintln(writer)
	}
}
