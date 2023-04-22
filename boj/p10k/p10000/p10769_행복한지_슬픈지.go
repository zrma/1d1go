package p10000

import (
	"bufio"
	"io"
)

func Solve10769(scanner *bufio.Scanner, writer io.Writer) {
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	s := scanner.Text()

	happy := 0
	sad := 0

	for i := 0; i < len(s)-2; i++ {
		if s[i] == ':' {
			switch s[i+1] {
			case '-':
				switch s[i+2] {
				case ')':
					happy++
				case '(':
					sad++
				}
			}
		}
	}

	if happy == 0 && sad == 0 {
		_, _ = writer.Write([]byte("none"))
	} else if happy == sad {
		_, _ = writer.Write([]byte("unsure"))
	} else if happy > sad {
		_, _ = writer.Write([]byte("happy"))
	} else {
		_, _ = writer.Write([]byte("sad"))
	}
}
