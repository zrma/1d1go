package p18100

import (
	"fmt"
	"io"
)

func Solve18108(reader io.Reader, writer io.Writer) {
	var year int
	_, _ = fmt.Fscan(reader, &year)

	res := year - 543
	_, _ = fmt.Fprint(writer, res)
}
