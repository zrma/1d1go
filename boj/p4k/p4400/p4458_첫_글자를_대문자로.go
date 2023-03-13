package p4400

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Solve4458(scanner *bufio.Scanner, writer io.Writer) {
	scanner.Split(bufio.ScanLines)

	scanner.Scan()

	var n int
	_, _ = fmt.Sscan(scanner.Text(), &n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()
		_, _ = fmt.Fprintf(writer, "%s%s\n", strings.ToUpper(string(s[0])), s[1:])
	}
}
