package p10800

import (
	"fmt"
	"strconv"
)

func Solve10869(scanner Scanner, writer Writer) {
	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())

	_, _ = fmt.Fprintln(writer, a+b)
	_, _ = fmt.Fprintln(writer, a-b)
	_, _ = fmt.Fprintln(writer, a*b)
	_, _ = fmt.Fprintln(writer, a/b)
	_, _ = fmt.Fprintln(writer, a%b)
}
