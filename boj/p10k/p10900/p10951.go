package p10900

import (
	"fmt"
	"strconv"
)

func Solve10951(scanner Scanner, writer Writer) {
	for scanner.Scan() {
		a, _ := strconv.Atoi(scanner.Text())
		if ok := scanner.Scan(); !ok {
			break
		}
		b, _ := strconv.Atoi(scanner.Text())

		_, _ = fmt.Fprintln(writer, a+b)
	}
	_ = writer.Flush()
}
