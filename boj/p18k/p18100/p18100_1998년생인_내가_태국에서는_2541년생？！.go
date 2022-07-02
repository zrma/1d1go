package p18100

import (
	"fmt"
)

func Solve18108(reader Reader, writer Writer) {
	var year int
	_, _ = fmt.Fscan(reader, &year)

	res := year - 543
	_, _ = fmt.Fprint(writer, res)
}
