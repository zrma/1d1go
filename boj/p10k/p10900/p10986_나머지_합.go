package p10900

import (
	"fmt"
)

func Solve10986(reader Reader, writer Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	arr := make([]int, m)
	var v int
	sum := 0
	res := 0
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(reader, &v)
		sum += v
		res += arr[sum%m]
		arr[sum%m]++
	}
	_, _ = fmt.Fprint(writer, res+arr[0])
}
