package p2600

import (
	"fmt"
	"strconv"
)

func Solve2675(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		r, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		for _, c := range scanner.Text() {
			for j := 0; j < r; j++ {
				_, _ = fmt.Fprint(writer, string(c))
			}
		}
		_, _ = fmt.Fprintln(writer)
	}
}
