package p11600

import (
	"fmt"
)

func Solve11654(scanner Scanner, writer Writer) {
	scanner.Scan()
	s := scanner.Text()
	_, _ = fmt.Fprint(writer, int(s[0]))
}
