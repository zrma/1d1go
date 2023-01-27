package p11700

import (
	"fmt"
	"io"

	"1d1go/utils/integer"
)

func Solve11758(reader io.Reader, writer io.Writer) {
	var x1, y1, x2, y2, x3, y3 int
	_, _ = fmt.Fscan(reader, &x1, &y1, &x2, &y2, &x3, &y3)

	res := integer.CCW(x1, y1, x2, y2, x3, y3)
	_, _ = fmt.Fprint(writer, res)
}
