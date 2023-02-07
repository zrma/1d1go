package p9000

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Solve9093(scanner *bufio.Scanner, writer io.Writer) {
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	var t int
	_, _ = fmt.Sscan(scanner.Text(), &t)

	for i := 0; i < t; i++ {
		scanner.Scan()
		input := scanner.Text()

		ss := strings.Split(input, " ")
		for j, s := range ss {
			if j < len(ss)-1 {
				_, _ = fmt.Fprint(writer, reverse(s), " ")
			} else {
				_, _ = fmt.Fprint(writer, reverse(s))
			}
		}
		_, _ = fmt.Fprintln(writer)
	}
}

func reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for i, r := range s {
		runes[n-1-i] = r
	}
	return string(runes)
}
