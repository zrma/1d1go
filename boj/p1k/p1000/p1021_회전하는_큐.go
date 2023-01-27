package p1000

import (
	"fmt"
	"io"
)

func Solve1021(reader io.Reader, writer io.Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	arr := make([]int, n)
	for i := range arr {
		arr[i] = i + 1
	}

	res := 0
	begin := 0
	for i := 0; i < m; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)

		length := len(arr)

		left := 0
		for v != arr[(begin+left)%length] && left < length {
			left++
		}
		right := 0
		for v != arr[(begin+length-right)%length] && right < length {
			right++
		}

		if left > right {
			res += right
			begin = (begin + length - right) % length
		} else {
			res += left
			begin = (begin + left) % length
		}
		arr = append(arr[:begin], arr[begin+1:]...)
	}
	_, _ = fmt.Fprint(writer, res)
}
