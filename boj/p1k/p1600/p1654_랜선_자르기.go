package p1600

import (
	"fmt"
	"sort"
)

func Solve1654(reader Reader, writer Writer) {
	var k, n int64
	_, _ = fmt.Fscan(reader, &k, &n)

	arr := make([]int, k)
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i])
	}
	sort.Ints(arr)

	var lo int64 = 1
	var hi = int64(arr[k-1])
	var res int64

	for lo <= hi {
		mid := lo + (hi-lo)/2

		var cnt int64
		for _, v := range arr {
			cnt += int64(v) / mid
		}

		if cnt < n {
			hi = mid - 1
		} else {
			lo = mid + 1
			res = mid
		}
	}

	_, _ = fmt.Fprint(writer, res)
}
