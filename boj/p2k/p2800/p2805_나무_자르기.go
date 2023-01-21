package p2800

import (
	"fmt"
	"sort"
)

func Solve2805(reader Reader, writer Writer) {
	var n, m int64
	_, _ = fmt.Fscan(reader, &n, &m)

	arr := make([]int, n)
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i])
	}
	sort.Ints(arr)

	var lo int64 = 0
	var hi = int64(arr[len(arr)-1])
	var res int64

	for lo <= hi {
		mid := lo + (hi-lo)/2

		var tot int64
		for _, v := range arr {
			cut := int64(v) - mid
			if cut < 0 {
				cut = 0
			}
			tot += cut
		}
		if tot < m {
			hi = mid - 1
		} else {
			lo = mid + 1
			res = mid
		}
	}

	_, _ = fmt.Fprint(writer, res)
}
