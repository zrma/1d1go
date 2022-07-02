package p7500

import (
	"fmt"
)

func Solve7568(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([][2]int, n)
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i][0], &arr[i][1])
	}

	for i, v0 := range arr {
		cnt := 1
		for j, v1 := range arr {
			if i == j {
				continue
			}
			if v0[0] < v1[0] && v0[1] < v1[1] {
				cnt++
			}
		}
		_, _ = fmt.Fprintf(writer, "%d ", cnt)
	}
}
