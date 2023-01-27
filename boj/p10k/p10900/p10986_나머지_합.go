package p10900

import (
	"fmt"
	"io"
)

func Solve10986(reader io.Reader, writer io.Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	arr := make([]int, m)
	sum := 0
	res := 0
	for i := 0; i < n; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)
		sum += v
		res += arr[sum%m]
		arr[sum%m]++
	}
	_, _ = fmt.Fprint(writer, res+arr[0])
}
