package p1800

import (
	"fmt"
)

func Solve1874(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	var ops []string
	stack := make([]int, 0, n)
	num := 0
	for i := 0; i < n; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)

		for j := num; j < v; j++ {
			num++
			stack = append(stack, num)
			ops = append(ops, "+")
		}
		if stack[len(stack)-1] == v {
			stack = stack[:len(stack)-1]
			ops = append(ops, "-")
		} else {
			_, _ = fmt.Fprint(writer, "NO")
			return
		}
	}
	for _, op := range ops {
		_, _ = fmt.Fprintln(writer, op)
	}
}
