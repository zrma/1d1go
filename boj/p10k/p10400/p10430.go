package p10400

import (
	"fmt"
	"strconv"
)

func Solve10430(scanner Scanner, writer Writer) {
	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	c, _ := strconv.Atoi(scanner.Text())

	_, _ = fmt.Fprintln(writer, (a+b)%c)
	_, _ = fmt.Fprintln(writer, (a%c+b%c)%c)
	_, _ = fmt.Fprintln(writer, (a*b)%c)
	_, _ = fmt.Fprintln(writer, (a%c*b%c)%c)
}
