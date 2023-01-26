package p11700

import (
	"bufio"
	"fmt"
)

func Solve11718(scanner *bufio.Scanner, writer Writer) {
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		s := scanner.Text()
		_, _ = fmt.Fprintln(writer, s)
	}
}
