package p1200

import (
	"fmt"
	"io"
	"math/big"
)

func Solve1247(reader io.Reader, writer io.Writer) {
	for i := 0; i < 3; i++ {
		var n int
		_, _ = fmt.Fscan(reader, &n)

		sum := big.NewInt(0)

		for j := 0; j < n; j++ {
			var x int64
			_, _ = fmt.Fscan(reader, &x)

			sum = sum.Add(sum, big.NewInt(x))
		}

		switch sum.Sign() {
		case 1:
			_, _ = fmt.Fprintln(writer, "+")
		case -1:
			_, _ = fmt.Fprintln(writer, "-")
		default:
			_, _ = fmt.Fprintln(writer, "0")
		}
	}
}
