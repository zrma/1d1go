package p11600

import (
	"fmt"
)

func Solve11654(reader Reader, writer Writer) {
	var s string
	_, _ = fmt.Fscan(reader, &s)
	_, _ = fmt.Fprint(writer, int(s[0]))
}
