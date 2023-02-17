package p10800

import (
	"bufio"
	"fmt"
	"io"
)

func Solve10820(scanner *bufio.Scanner, writer io.Writer) {
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Bytes()
		var lower, upper, number, space int
		for _, c := range line {
			switch {
			case c >= 'a' && c <= 'z':
				lower++
			case c >= 'A' && c <= 'Z':
				upper++
			case c >= '0' && c <= '9':
				number++
			case c == ' ':
				space++
			}
		}
		_, _ = fmt.Fprintf(writer, "%d %d %d %d\n", lower, upper, number, space)
	}
}
