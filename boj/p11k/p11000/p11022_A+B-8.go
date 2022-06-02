package p11000

import (
	"fmt"
	"strconv"
)

func Solve11022(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())

		_, _ = fmt.Fprintf(writer, "Case #%d: %d + %d = %d\n", i+1, a, b, a+b)
	}
}
