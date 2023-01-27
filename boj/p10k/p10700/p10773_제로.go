package p10700

import (
	"fmt"
	"io"
)

func Solve10773(reader io.Reader, writer io.Writer) {
	var k int
	_, _ = fmt.Fscan(reader, &k)

	arr := make([]int, 0, k)
	for i := 0; i < k; i++ {
		var n int
		_, _ = fmt.Fscan(reader, &n)
		if n == 0 {
			arr = arr[:len(arr)-1]
			continue
		}
		arr = append(arr, n)
	}

	sum := 0
	for _, n := range arr {
		sum += n
	}
	_, _ = fmt.Fprint(writer, sum)
}
