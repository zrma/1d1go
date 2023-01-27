package p11000

import (
	"fmt"
	"io"

	"1d1go/utils/integer"
)

func Solve11051(reader io.Reader, writer io.Writer) {
	var n, k int
	_, _ = fmt.Fscan(reader, &n, &k)

	const maxLen = 1000 + 1
	cache := integer.BuildCache2DArr(maxLen, -1)

	res := integer.CombinationDP(n, k, cache, func(v int) int {
		return v % 10_007
	})
	_, _ = fmt.Fprint(writer, res)
}
