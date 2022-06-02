package p11700

import (
	"fmt"
	"strconv"
)

func Solve11720(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	s := scanner.Text()

	var sum int
	for i := 0; i < n && i < len(s); i++ {
		c := int(s[i] - '0')
		sum += c
	}
	_, _ = fmt.Fprintf(writer, "%d", sum)
}
