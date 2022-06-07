package p11000

import (
	"fmt"
	"strconv"

	"1d1go/utils/integer"
)

func Solve11051(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	const maxLen = 1000 + 1
	cache := integer.BuildCache2DArr(maxLen, -1)

	res := integer.CombinationDP(n, k, cache, func(v int) int {
		return v % 10_007
	})
	_, _ = fmt.Fprint(writer, res)
}
