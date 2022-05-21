package p1400

import (
	"fmt"
	"strconv"

	"1d1go/utils/integer"
)

func Solve1463(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n+1)
	for i := 2; i <= n; i++ {
		arr[i] = arr[i-1] + 1
		if i%2 == 0 {
			arr[i] = integer.MinInt(arr[i], arr[i/2]+1)
		}
		if i%3 == 0 {
			arr[i] = integer.MinInt(arr[i], arr[i/3]+1)
		}
	}

	_, _ = fmt.Fprint(writer, arr[n])
}
