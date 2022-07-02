package p2500

import (
	"fmt"
)

func Solve2588(reader Reader, writer Writer) {
	var a, b int
	_, _ = fmt.Fscan(reader, &a, &b)

	third := a * (b % 10)
	fourth := a * (b / 10 % 10)
	fifth := a * (b / 100 % 10)

	_, _ = fmt.Fprintln(writer, third)
	_, _ = fmt.Fprintln(writer, fourth)
	_, _ = fmt.Fprintln(writer, fifth)
	_, _ = fmt.Fprintln(writer, third+fourth*10+fifth*100)
}
