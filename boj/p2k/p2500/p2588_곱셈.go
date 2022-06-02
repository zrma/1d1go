package p2500

import (
	"fmt"
	"strconv"
)

func Solve2588(scanner Scanner, writer Writer) {
	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())

	third := a * (b % 10)
	fourth := a * (b / 10 % 10)
	fifth := a * (b / 100 % 10)

	_, _ = fmt.Fprintln(writer, third)
	_, _ = fmt.Fprintln(writer, fourth)
	_, _ = fmt.Fprintln(writer, fifth)
	_, _ = fmt.Fprintln(writer, third+fourth*10+fifth*100)
}
