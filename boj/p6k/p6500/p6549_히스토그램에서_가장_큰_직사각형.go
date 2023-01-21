package p6500

import (
	"fmt"

	"1d1go/utils/integer"
)

func Solve6549(reader Reader, writer Writer) {
	for {
		var n int64
		_, _ = fmt.Fscan(reader, &n)

		if n == 0 {
			break
		}

		heights = make([]int64, n)
		for i := range heights {
			_, _ = fmt.Fscan(reader, &heights[i])
		}

		res := findMaxArea(0, len(heights)-1)
		_, _ = fmt.Fprintln(writer, res)
	}
}

var heights []int64

func findMaxArea(left, right int) int64 {
	if left == right {
		return heights[left]
	}
	mid := left + (right-left)/2
	res := integer.Max(findMaxArea(left, mid), findMaxArea(mid+1, right))

	lo := mid
	hi := mid + 1

	height := integer.Min(heights[lo], heights[hi])
	res = integer.Max(res, height*2)

	for left < lo || hi < right {
		if hi < right && (lo == left || heights[lo-1] < heights[hi+1]) {
			hi++
			height = integer.Min(height, heights[hi])
		} else {
			lo--
			height = integer.Min(height, heights[lo])
		}
		res = integer.Max(res, height*int64(hi-lo+1))
	}
	return res
}
