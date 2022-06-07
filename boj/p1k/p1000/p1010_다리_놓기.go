package p1000

import (
	"fmt"
	"strconv"

	"1d1go/utils/integer"
)

func Solve1010(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	const maxLen = 30
	cache := integer.BuildCache2DArr(maxLen, -1)

	for i := 0; i < n; i++ {
		scanner.Scan()
		from, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		to, _ := strconv.Atoi(scanner.Text())

		res := integer.CombinationDP(to, from, cache, func(v int) int {
			return v
		})
		_, _ = fmt.Fprintln(writer, res)
	}
}
