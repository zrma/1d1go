package p1600

import (
	"fmt"
	"sort"
)

func Solve1654(reader Reader, writer Writer) {
	var k, n int
	_, _ = fmt.Fscan(reader, &k, &n)

	arr := make([]int, k)
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i])
	}
	sort.Ints(arr)
}
