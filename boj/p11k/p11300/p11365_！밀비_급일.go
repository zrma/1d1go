package p11300

import (
	"bufio"
	"io"
)

func Solve11365(scanner *bufio.Scanner, writer io.Writer) {
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		s := scanner.Text()
		if s == "END" {
			break
		}

		for i := len(s) - 1; i >= 0; i-- {
			_, _ = writer.Write([]byte{s[i]})
		}
		_, _ = writer.Write([]byte{'\n'})
	}
}
