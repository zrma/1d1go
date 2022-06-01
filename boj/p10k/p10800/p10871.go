package p10800

import (
	"fmt"
	"strconv"
)

func Solve10871(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		if v < x {
			_, _ = fmt.Fprintf(writer, "%d ", v)
		}
	}
}
