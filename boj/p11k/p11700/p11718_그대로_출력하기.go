package p11700

import (
	"bufio"
	"fmt"
	"io"
)

func Solve11718(scanner *bufio.Scanner, writer io.Writer) {
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		s := scanner.Text()
		_, _ = fmt.Fprintln(writer, s)
	}
}
