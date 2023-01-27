package p1000

import (
	"fmt"
	"io"
	"sort"
)

func Solve1015(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([]int, n)
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	sortedArr := make([]int, n)
	copy(sortedArr, arr)
	sort.Ints(sortedArr)

	visited := make([]bool, n)
	for i, before := range arr {
		for j, after := range sortedArr {
			if !visited[j] && before == after {
				if i < n-1 {
					_, _ = fmt.Fprint(writer, j, " ")
				} else {
					_, _ = fmt.Fprint(writer, j)
				}
				visited[j] = true
				break
			}
		}
	}
}
