package p1300

import (
	"bufio"
	"io"
)

func Solve1371(scanner *bufio.Scanner, writer io.Writer) {
	scanner.Split(bufio.ScanLines)

	counts := make([]int, 26)
	for scanner.Scan() {
		s := scanner.Text()
		for _, c := range s {
			if c >= 'a' && c <= 'z' {
				counts[c-'a']++
			}
		}
	}

	max := 0
	for _, c := range counts {
		if c > max {
			max = c
		}
	}

	for i, c := range counts {
		if c == max {
			_, _ = writer.Write([]byte{byte(i) + 'a'})
		}
	}
}
