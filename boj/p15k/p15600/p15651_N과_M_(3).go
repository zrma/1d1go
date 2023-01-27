package p15600

import (
	"fmt"
	"io"
)

func Solve15651(reader io.Reader, writer io.Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	arr := make([]int, m)
	permutationWithRepetitionDFS15651(arr, 0, n, m, writer)
}

func permutationWithRepetitionDFS15651(arr []int, depth, n, m int, writer io.Writer) {
	if depth >= m {
		for i, v := range arr {
			_, _ = fmt.Fprint(writer, v)
			if i != m-1 {
				_, _ = fmt.Fprint(writer, " ")
			}
		}
		_, _ = fmt.Fprintln(writer)
		return
	}

	for i := 0; i < n; i++ {
		arr[depth] = i + 1
		permutationWithRepetitionDFS15651(arr, depth+1, n, m, writer)
	}
}
