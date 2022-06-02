package p8900

import (
	"fmt"
	"strconv"
)

func Solve8958(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()

		res := acc(s)
		_, _ = fmt.Fprintln(writer, res)
	}
}

func acc(s string) int {
	tot := 0
	cur := 1
	for _, c := range s {
		if c == 'O' {
			tot += cur
			cur++
		} else {
			cur = 1
		}
	}
	return tot
}
