package p1800

import (
	"fmt"
)

func Solve1806(reader Reader, writer Writer) {
	var n, s int
	_, _ = fmt.Fscan(reader, &n, &s)

	arr := make([]int, n+1)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	begin, end := 0, 0
	sum := arr[0]
	min := n + 1

	for begin <= end && end < n {
		if sum < s {
			end++
			sum += arr[end]
		} else {
			if min > end-begin+1 {
				min = end - begin + 1
			}
			sum -= arr[begin]
			begin++
		}
	}

	if min > n {
		_, _ = fmt.Fprint(writer, 0)
	} else {
		_, _ = fmt.Fprint(writer, min)
	}
}
