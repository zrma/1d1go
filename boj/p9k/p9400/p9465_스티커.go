package p9400

import (
	"fmt"

	"1d1go/utils/integer"
)

func Solve9465(reader Reader, writer Writer) {
	var t int
	_, _ = fmt.Fscan(reader, &t)

	for ; t > 0; t-- {
		var n int
		_, _ = fmt.Fscan(reader, &n)

		arr := make([][2]int, n+2)
		for i := 2; i < n+2; i++ {
			_, _ = fmt.Fscan(reader, &arr[i][0])
		}
		for i := 2; i < n+2; i++ {
			_, _ = fmt.Fscan(reader, &arr[i][1])
		}

		_, _ = fmt.Fprintln(writer, solve9465(arr, n))
	}
}

func solve9465(arr [][2]int, n int) int {
	for i := 2; i < n+2; i++ {
		arr[i][0] += integer.Max(arr[i-1][1], arr[i-2][1])
		arr[i][1] += integer.Max(arr[i-1][0], arr[i-2][0])
	}

	return integer.Max(arr[n+1][0], arr[n+1][1])
}
