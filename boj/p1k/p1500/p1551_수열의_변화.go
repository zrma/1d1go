package p1500

import (
	"fmt"
	"io"
)

func Solve1551(reader io.Reader, at io.Writer) {
	var n, k int
	_, _ = fmt.Fscanln(reader, &n, &k)

	arr := make([]int, n)
	for i := range arr {
		_, _ = fmt.Fscanf(reader, "%d,", &arr[i])
	}

	for i := 0; i < k; i++ {
		for j := 0; j < len(arr)-1; j++ {
			arr[j] = arr[j+1] - arr[j]
		}
		arr = arr[:len(arr)-1]
	}

	for i, v := range arr {
		if i == len(arr)-1 {
			_, _ = fmt.Fprintf(at, "%d", v)
		} else {
			_, _ = fmt.Fprintf(at, "%d,", v)
		}
	}
}
