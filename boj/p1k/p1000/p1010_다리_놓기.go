package p1000

import (
	"fmt"

	"1d1go/utils/integer"
)

func Solve1010(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	const maxLen = 30
	cache := integer.BuildCache2DArr(maxLen, -1)

	for i := 0; i < n; i++ {
		var from, to int
		_, _ = fmt.Fscan(reader, &from, &to)

		res := integer.CombinationDP(to, from, cache, func(v int) int {
			return v
		})
		_, _ = fmt.Fprintln(writer, res)
	}
}
