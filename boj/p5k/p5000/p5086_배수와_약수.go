package p5000

import (
	"fmt"
	"strconv"
)

func Solve5086(scanner Scanner, writer Writer) {
	for {
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())

		if a == 0 && b == 0 {
			break
		}

		if b%a == 0 {
			_, _ = fmt.Fprintln(writer, "factor")
		} else if a%b == 0 {
			_, _ = fmt.Fprintln(writer, "multiple")
		} else {
			_, _ = fmt.Fprintln(writer, "neither")
		}
	}
}
