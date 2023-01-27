package p1000

import (
	"fmt"
	"io"
)

func Solve1024(reader io.Reader, writer io.Writer) {
	var n, l int
	_, _ = fmt.Fscan(reader, &n, &l)

	// n = (x + 0) + (x + 1) + ... + (x + l - 1)
	// n = l * x + (0 + 2 + ... + l - 1)
	// n = l * x + (l - 1) * l / 2
	// n - (l - 1) * l / 2 = l * x

	for i := l; i <= 100; i++ {
		lx := n - (i-1)*i/2

		if lx%i != 0 {
			continue
		}

		x := lx / i
		if x >= 0 {
			for j := 0; j < i; j++ {
				_, _ = fmt.Fprint(writer, x+j)
				if j != i-1 {
					_, _ = fmt.Fprint(writer, " ")
				}
			}
			return
		}
	}

	_, _ = fmt.Fprint(writer, -1)
}
