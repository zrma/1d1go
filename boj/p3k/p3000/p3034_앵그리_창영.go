package p3000

import (
	"fmt"
	"strconv"
)

func Solve3034(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	w, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	h, _ := strconv.Atoi(scanner.Text())

	limit := w*w + h*h

	for i := 0; i < n; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())

		if v*v <= limit {
			_, _ = fmt.Fprintln(writer, "DA")
		} else {
			_, _ = fmt.Fprintln(writer, "NE")
		}
	}
}
