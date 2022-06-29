package p9300

import (
	"fmt"

	"1d1go/programmers/practice"
)

func Solve9375(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		var k int
		_, _ = fmt.Fscan(reader, &k)

		arr := make([][]string, k)
		for j := range arr {
			arr[j] = make([]string, 2)
			_, _ = fmt.Fscan(reader, &arr[j][0], &arr[j][1])
		}
		res := practice.Disguise(arr)
		_, _ = fmt.Fprintln(writer, res)
	}
}
