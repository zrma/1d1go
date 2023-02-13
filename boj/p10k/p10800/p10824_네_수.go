package p10800

import (
	"fmt"
	"io"
)

func Solve10824(reader io.Reader, writer io.Writer) {
	var a, b, c, d int
	_, _ = fmt.Fscan(reader, &a, &b, &c, &d)

	decimal := b
	for decimal > 0 {
		decimal /= 10
		a *= 10
	}

	decimal = d
	for decimal > 0 {
		decimal /= 10
		c *= 10
	}

	_, _ = fmt.Fprint(writer, a+b+c+d)
}
