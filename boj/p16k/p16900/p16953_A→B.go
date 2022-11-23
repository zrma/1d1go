package p16900

import (
	"fmt"
)

func Solve16953(reader Reader, writer Writer) {
	var a, b int
	_, _ = fmt.Fscan(reader, &a, &b)

	res := 0
	for b > a {
		if b%2 == 0 {
			b /= 2
		} else if b%10 == 1 {
			b /= 10
		} else {
			break
		}
		res++
	}

	if a == b {
		_, _ = fmt.Fprint(writer, res+1)
	} else {
		_, _ = fmt.Fprint(writer, -1)
	}
}
