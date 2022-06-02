package p4100

import (
	"fmt"
	"strconv"
)

func Solve4153(scanner Scanner, writer Writer) {
	for scanner.Scan() {
		a, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		c, _ := strconv.Atoi(scanner.Text())

		if a == 0 && b == 0 && c == 0 {
			return
		}

		if a*a+b*b == c*c || a*a+c*c == b*b || b*b+c*c == a*a {
			_, _ = fmt.Fprintln(writer, "right")
		} else {
			_, _ = fmt.Fprintln(writer, "wrong")
		}
	}
}
