package p9400

import (
	"fmt"
	"strconv"
)

func Solve9461(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	const maxLen = 100 + 1
	arr := make([]int, maxLen)

	arr[0] = 0
	arr[1] = 1
	arr[2] = 1
	arr[3] = 1
	arr[4] = 2

	for i := 5; i < maxLen; i++ {
		arr[i] = arr[i-1] + arr[i-5]
	}

	for i := 0; i < n; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())

		res := arr[v]
		_, _ = fmt.Fprintln(writer, res)
	}
}
