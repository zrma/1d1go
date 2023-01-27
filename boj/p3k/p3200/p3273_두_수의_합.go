package p3200

import (
	"fmt"
	"io"
	"sort"
)

func Solve3273(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	var x int
	_, _ = fmt.Fscan(reader, &x)

	const max = 2_000_000
	counts := make([]int, max+1)

	res := 0
	for _, v := range arr {
		if v >= x {
			continue
		}

		if counts[v] > 0 {
			res++
			counts[v]--
			continue
		}

		counts[x-v]++
	}

	_, _ = fmt.Fprint(writer, res)
}

func Solve3273WithTwoPoints(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	var x int
	_, _ = fmt.Fscan(reader, &x)

	sort.Ints(arr)

	res := 0
	lo, hi := 0, n-1
	for lo < hi {
		if arr[lo]+arr[hi] == x {
			res++
			lo++
			hi--
			continue
		}

		if arr[lo]+arr[hi] < x {
			lo++
			continue
		}

		hi--
	}

	_, _ = fmt.Fprint(writer, res)
}
