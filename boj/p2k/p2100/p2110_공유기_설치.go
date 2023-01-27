package p2100

import (
	"fmt"
	"io"
	"sort"
)

func Solve2110(reader io.Reader, writer io.Writer) {
	var n, c int64
	_, _ = fmt.Fscan(reader, &n, &c)

	arr := make([]int64, n)
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	lo := int64(1)
	hi := arr[n-1] - arr[0]
	var res int64

	for lo <= hi {
		mid := lo + (hi-lo)/2

		cnt := int64(1)
		prev := arr[0]
		for i := int64(1); i < n; i++ {
			if dist := arr[i] - prev; dist < mid {
				continue
			}
			cnt++
			prev = arr[i]
		}

		if cnt < c {
			hi = mid - 1
		} else {
			lo = mid + 1
			res = mid
		}
	}

	_, _ = fmt.Fprint(writer, res)
}
