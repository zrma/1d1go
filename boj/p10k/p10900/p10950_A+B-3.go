package p10900

import (
	"fmt"
	"strconv"
)

func Solve10950(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())

		_, _ = fmt.Fprintln(writer, a+b)
	}
}
