package p10250

import (
	"fmt"
	"strconv"
)

func Solve10250(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		h, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		w, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())

		res := calcNearestRoom(h, w, k)
		_, _ = fmt.Fprintln(writer, res)
	}
}

func calcNearestRoom(h, _, n int) int {
	if n == 1 {
		return 101
	}
	col := n/h + 1
	row := n % h
	if row == 0 {
		row = h
		col -= 1
	}
	return row*100 + col
}
