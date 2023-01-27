package p17200

import (
	"fmt"
	"io"
)

func Solve17299(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	const max = 1_000_000 + 1

	arr := make([]int, n)
	res := make([]int, n)
	counts := make([]int, max)

	for i := 0; i < n; i++ {
		res[i] = -1
		_, _ = fmt.Fscan(reader, &arr[i])
		counts[arr[i]]++
	}

	stack := make([]int, 0, n)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && counts[arr[stack[len(stack)-1]]] < counts[arr[i]] {
			res[stack[len(stack)-1]] = arr[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	for i := 0; i < n; i++ {
		_, _ = fmt.Fprintf(writer, "%d ", res[i])
	}
}
