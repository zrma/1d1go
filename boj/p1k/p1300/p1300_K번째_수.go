package p1300

import (
	"fmt"
	"io"
)

func Solve1300(reader io.Reader, writer io.Writer) {
	var n, k int64
	_, _ = fmt.Fscan(reader, &n, &k)

	lo := int64(1)
	hi := n * n
	res := int64(0)

	for lo <= hi {
		mid := lo + (hi-lo)/2

		cnt := int64(0)
		for i := int64(1); i <= n; i++ {
			v := mid / i
			if v > n {
				v = n
			}
			cnt += v
		}
		if cnt < k {
			lo = mid + 1
		} else {
			hi = mid - 1
			res = mid
		}
	}

	_, _ = fmt.Fprint(writer, res)
}
