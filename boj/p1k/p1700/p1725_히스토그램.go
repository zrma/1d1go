package p1700

import (
	"fmt"
	"io"
)

func Solve1725(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([]int, n+1)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	stack := make([]int, 0, n)
	stack = append(stack, 0)
	res := 0
	for i := 1; i <= n; i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > arr[i] {
			h := arr[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			w := i
			if len(stack) > 0 {
				w = i - stack[len(stack)-1] - 1
			}

			if res < h*w {
				res = h * w
			}
		}
		stack = append(stack, i)
	}

	_, _ = fmt.Fprint(writer, res)
}
