package p10800

import (
	"fmt"
	"io"
)

func Solve10807(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([]int, 200+2)
	for i := 0; i < n; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)
		arr[v+100]++
	}

	var v int
	_, _ = fmt.Fscan(reader, &v)
	_, _ = fmt.Fprint(writer, arr[v+100])
}
