package p2800

import (
	"fmt"
	"strconv"

	"1d1go/utils/integer"
)

func Solve2839(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arrSize := n + 6
	if arrSize < 6 {
		arrSize = 6
	}
	arr := make([]int, arrSize)
	for i := range arr {
		arr[i] = -1
	}
	arr[3] = 1
	arr[5] = 1

	for i := 3; i <= n; i++ {
		if arr[i] == -1 {
			continue
		}
		if arr[i+3] == -1 {
			arr[i+3] = arr[i] + 1
		} else {
			arr[i+3] = integer.MinInt(arr[i+3], arr[i]+1)
		}
		if arr[i+5] == -1 {
			arr[i+5] = arr[i] + 1
		} else {
			arr[i+5] = integer.MinInt(arr[i+5], arr[i]+1)
		}
	}
	_, _ = fmt.Fprint(writer, arr[n])
}
