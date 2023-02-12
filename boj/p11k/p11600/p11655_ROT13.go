package p11600

import (
	"bufio"
	"fmt"
	"io"
)

func Solve11655(scanner *bufio.Scanner, writer io.Writer) {
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	s := scanner.Text()

	for _, c := range s {
		if c >= 'a' && c <= 'm' || c >= 'A' && c <= 'M' {
			_, _ = fmt.Fprint(writer, string(c+13))
		} else if c >= 'n' && c <= 'z' || c >= 'N' && c <= 'Z' {
			_, _ = fmt.Fprint(writer, string(c-13))
		} else {
			_, _ = fmt.Fprint(writer, string(c))
		}
	}
}
