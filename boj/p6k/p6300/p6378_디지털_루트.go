package p6300

import (
	"bufio"
	"fmt"
	"io"
)

func Solve6378(scanner *bufio.Scanner, writer io.Writer) {
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		s := scanner.Text()
		if s == "0" {
			break
		}

		sum := 0
		for _, c := range s {
			sum += int(c - '0')
		}

		for sum >= 10 {
			sum = sum/10 + sum%10
		}

		_, _ = fmt.Fprintln(writer, sum)
	}
}
