package p3000

import (
	"fmt"
	"io"
)

func Solve3015(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	stack := make([][2]int, 0, n)

	res := 0
	for i := 0; i < n; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)

		for len(stack) > 0 && stack[len(stack)-1][0] < v {
			res += stack[len(stack)-1][1]
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			stack = append(stack, [2]int{v, 1})
			continue
		}

		if stack[len(stack)-1][0] == v {
			res += stack[len(stack)-1][1]
			stack[len(stack)-1][1]++

			if len(stack) > 1 {
				res++
			}
			continue
		}

		res++
		stack = append(stack, [2]int{v, 1})
	}

	_, _ = fmt.Fprint(writer, res)
}
