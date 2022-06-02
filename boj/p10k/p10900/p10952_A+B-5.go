package p10900

import (
	"fmt"
	"strconv"
)

func Solve10952(scanner Scanner, writer Writer) {
	for scanner.Scan() {
		a, _ := strconv.Atoi(scanner.Text())
		if ok := scanner.Scan(); !ok {
			break
		}
		b, _ := strconv.Atoi(scanner.Text())

		if a == 0 && b == 0 {
			break
		}

		_, _ = fmt.Fprintln(writer, a+b)
	}
	_ = writer.Flush()
}
