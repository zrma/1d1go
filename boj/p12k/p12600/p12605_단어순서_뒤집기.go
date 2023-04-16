package p12600

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Solve12605(scanner *bufio.Scanner, writer io.Writer) {
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		_, _ = fmt.Fprintf(writer, "Case #%d: ", i+1)

		scanner.Scan()
		line := scanner.Text()

		words := strings.Split(line, " ")
		for j := len(words) - 1; j >= 0; j-- {
			_, _ = fmt.Fprint(writer, words[j])
			if j > 0 {
				_, _ = fmt.Fprint(writer, " ")
			}
		}
		_, _ = fmt.Fprintln(writer)
	}
}
