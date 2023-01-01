package p1500

import (
	"bufio"
	"fmt"
)

func Solve1543(scanner *bufio.Scanner, writer Writer) {
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	s := scanner.Text()
	scanner.Scan()
	t := scanner.Text()

	_, _ = fmt.Fprint(writer, solve1543(s, t))
}

func solve1543(s string, t string) int {
	res := 0
	idx := 0
	for idx < len(s) {
		if idx+len(t) > len(s) {
			break
		}

		if s[idx:idx+len(t)] == t {
			res++
			idx += len(t)
		} else {
			idx++
		}
	}
	return res
}
