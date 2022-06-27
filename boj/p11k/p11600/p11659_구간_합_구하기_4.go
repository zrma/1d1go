package p11600

import (
	"fmt"
	"strconv"
)

func Solve11659(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n+1)

	for i := 0; i < n; i++ {
		scanner.Scan()
		arr[i+1], _ = strconv.Atoi(scanner.Text())
		arr[i+1] += arr[i]
	}

	for i := 0; i < m; i++ {
		scanner.Scan()
		begin, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		end, _ := strconv.Atoi(scanner.Text())

		res := arr[end] - arr[begin-1]
		_, _ = fmt.Fprintln(writer, res)
	}
}
